package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/users"
	"os"
)

type resourceLister struct{}
type userObjectExistsChecker struct{}

var resourceList resourceLister
var userObjectExistsCheck userObjectExistsChecker

func init() {
	resourceList = resourceLister{}
}

func (resourceList *resourceLister) ListItems(partitionKey interface{}, keyPrefix interface{}) (interface{}, error) {
	fmt.Printf("The partion key: %v, the sort key prefix: %v", partitionKey, keyPrefix)
	return dbUtils.ListResourcesForObject(os.Getenv("TABLE_NAME"), partitionKey.(string), keyPrefix.(string))
}

func (userObjectExistsCheck *userObjectExistsChecker) CheckItemExists(partitionKey interface{}, rangeKey interface{}) (bool, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	return dbHelper.CheckItemExists(os.Getenv("TABLE_NAME"), partitionKey.(string), rangeKey.(string))
}

func handleResourceList(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	user := users.User{}
	user.Email, user.UserId = commons.GetUserProfileInfo(request)

	fmt.Printf("The request body: %v\n", request.Body)

	var object objects.Object
	err := commons.UnmarshallRequestBody(request.Body, &object)
	if err != nil {
		return commons.BuildResponse(400, struct {
			Error string `json:"error"`
		}{"Invalid Input."})
	}

	exists, err := userObjectExistsCheck.CheckItemExists(users.GetUserPK(user.Email), users.GetUserObjectSK(object.UOI))
	if err != nil {
		fmt.Printf("An error occurred while checking the userObj: %v/%v existence. Error: %v ", user.Email, object.UOI, err)
		return commons.BuildResponse(500, struct {
			Error string `json:"error"`
		}{"An error occurred we investigate why this happened."})
	}

	if !exists {
		return commons.BuildResponse(404, struct {
			Error string `json:"error"`
		}{"Resource not found."})
	}

	results, err := resourceList.ListItems(objects.GetObjectPK(object.UOI), "RSC#")
	if err != nil {
		return commons.BuildResponse(500, struct {
			Error string `json:"error"`
		}{"Invalid Input."})
	}

	return commons.BuildResponse(200, results)
}

func main() {
	lambda.Start(handleResourceList)
}
