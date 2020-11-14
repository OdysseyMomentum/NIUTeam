package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/resources"
	"github.com/srbry/go-serverless-example/commons/users"
	//TODO: Use mapstructure for the case of Resource Meta
	//"github.com/mitchellh/mapstructure"
	"os"
)

type telemetryLister struct{}
type userObjectExistsChecker struct{}

var telemetryList telemetryLister
var userObjectExistsCheck userObjectExistsChecker

func init() {
	telemetryList = telemetryLister{}
	userObjectExistsCheck = userObjectExistsChecker{}
}

//TODO move to a separate function for better unit testing.
func (telemetryList *telemetryLister) ListItems(partitionKey interface{}, keyPrefix interface{}) (interface{}, error) {
	fmt.Printf("The partion key: %v, the sort key prefix: %v", partitionKey, keyPrefix)
	return dbUtils.GetSensorData(os.Getenv("IOT_TABLE"), partitionKey.(string), keyPrefix.(resources.Range).Begin, keyPrefix.(resources.Range).End)
}

func (userObjectExistsCheck *userObjectExistsChecker) CheckItemExists(partitionKey interface{}, sortKey interface{}) (bool, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	fmt.Printf("OS Table Name: %v\n",os.Getenv("TABLE_NAME"))
	return dbHelper.CheckItemExists(os.Getenv("TABLE_NAME"), partitionKey.(string), sortKey.(string))
}

func handleSensorDataGet(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	email, _ := commons.GetUserProfileInfo(request)
	fmt.Printf("The Body: %v\n", request.Body)
	fmt.Printf("The Mail: %v\n", email)

	inputStruct := struct {
		Object   objects.Object     `json:"object"`
		Resource resources.Resource `json:"resource"`
		Range    resources.Range    `json:"range"`
	}{}
	err := commons.UnmarshallRequestBody(request.Body, &inputStruct)
	fmt.Printf("The inputStruct: %v\n", inputStruct)
	if err != nil {
		return commons.BuildResponse(400, "The input format is not correct." + err.Error())
	}

	exists, err := userObjectExistsCheck.CheckItemExists(users.GetUserPK(email), users.GetUserObjectSK(inputStruct.Object.UOI))
	if err != nil {
		return commons.BuildResponse(400, "The input format is not correct." + err.Error())
	}

	if !exists {
		return commons.BuildResponse(404, "The Resource is not found.")
	}

	items, err := telemetryList.ListItems(inputStruct.Resource.ResourceId, inputStruct.Range)
	if err != nil {
		fmt.Printf("An Error occured: %v\n", err)
		return commons.BuildResponse(404, "The Resource is not found.")
	}
	return commons.BuildResponse(200, items)
}

func main() {
	lambda.Start(handleSensorDataGet)
}
