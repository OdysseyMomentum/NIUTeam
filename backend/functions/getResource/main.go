package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/resources"
	"github.com/srbry/go-serverless-example/commons/users"
	//TODO: Use mapstructure for the case of Resource Meta
    //"github.com/mitchellh/mapstructure"
    "os"
	"time"
)

type resourceGetter struct{}
type userObjectExistsChecker struct{}

var resourceGet resourceGetter
var userObjectExistsCheck userObjectExistsChecker

func init() {
	resourceGet = resourceGetter{}
	userObjectExistsCheck = userObjectExistsChecker{}
}

func (userObjectExistsCheck *userObjectExistsChecker) CheckItemExists(partitionKey interface{}, sortKey interface{}) (bool, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	return dbHelper.CheckItemExists(os.Getenv("TABLE_NAME"), partitionKey.(string), sortKey.(string))
}

func (resourceGet *resourceGetter) GetItem(tableName string, partitionKey string, sortKey string) (interface{}, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	return dbHelper.GetItem(os.Getenv("TABLE_NAME"), partitionKey, sortKey)
}

func handleResourceGet(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	email, _ := commons.GetUserProfileInfo(request)
	fmt.Printf("The Body: %v\n", request.Body)
	fmt.Printf("The Mail: %v\n", email)

	inputStruct := struct {
		Object   objects.Object     `json:"object"`
		Resource resources.Resource `json:"resource"`
	}{}
	err := commons.UnmarshallRequestBody(request.Body, &inputStruct)
	fmt.Printf("The inputStruct: %v\n", inputStruct)
	if err != nil {
		return commons.BuildResponse(400, "The input format is not correct.")
	}

	exists, err := userObjectExistsCheck.CheckItemExists(users.GetUserPK(email), users.GetUserObjectSK(inputStruct.Object.UOI))
	if err != nil {
		return commons.BuildResponse(400, "The input format is not correct.")
	}

	if !exists {
		return commons.BuildResponse(404, "The Resource is not found.")
	}

	item, err := resourceGet.GetItem(os.Getenv("TABLE_NAME"), resources.GetResourcePK(inputStruct.Resource.ResourceId), "DETAILS")
	if err != nil {
		fmt.Printf("An Error occured: %v\n", err)
		return commons.BuildResponse(404, "The Resource is not found.")
	}

	resource := resources.Resource{}
	err = dynamodbattribute.UnmarshalMap(item.(map[string]*dynamodb.AttributeValue), &resource)

	if resource.ResourceType == "DOCUMENT" {

		fmt.Printf("The Resource: %+v\n", resource)

		key := fmt.Sprintf("%v/%v", inputStruct.Object.UOI, resource.ResourceId)
		fmt.Printf("The Key: %v\n", key)
		fmt.Printf("The Bucket: %v\n", os.Getenv("BUCKET_NAME"))
		fmt.Printf("The Filename: %v\n", resource.Meta.(map[string]interface{})["filename"].(string))

		//Get the DownloadURL
		svc := s3iface.S3API(s3.New(session.Must(session.NewSession())))
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket:                     aws.String(os.Getenv("BUCKET_NAME")),
			Key:                        aws.String(fmt.Sprintf("%v/%v", inputStruct.Object.UOI, resource.ResourceId)),
			ResponseContentDisposition: aws.String(fmt.Sprintf("filename=\"%v\"", resource.Meta.(map[string]interface{})["filename"].(string))),
			ResponseContentType: 		aws.String(commons.GetContentTypeForFile(resource.Meta.(map[string]interface{})["filetype"].(string))),
		})
		url, err := req.Presign(3 * time.Minute)
		if err != nil {
			fmt.Printf("An Error occured while presigning: %v\n", err)
			return commons.BuildResponse(500, "We have to investigate.")
		}
		//TODO see how we can convert the string map to a struct
		resource.Meta.(map[string]interface{})["access"] = url
	}

	if err != nil {
		fmt.Printf("An Error occured: %v\n", err)
		return commons.BuildResponse(404, "The Resource is not found.")
	}
	return commons.BuildResponse(200, resource)
}

func main() {
	lambda.Start(handleResourceGet)
}
