package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/objects"
	"os"
)

type objectGetter struct{}

var objectGet objectGetter

func init() {
	objectGet = objectGetter{}
}

func (objectGet *objectGetter) GetItem(partitionKey interface{}, rangeKey interface{}) (interface{}, error) {
	return dbUtils.GetObjectById(os.Getenv("TABLE_NAME"), objects.GetObjectPK(partitionKey.(string)))
}

func handleGetObject(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//TODO: See if the user is allowed to get them

	obj := objects.Object{}
	commons.UnmarshallRequestBody(request.Body, &obj)
	item, err := objectGet.GetItem(obj.UOI, "DETAILS")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	resp, err := commons.BuildResponse(200, item)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	return resp, nil
}

func main() {
	lambda.Start(handleGetObject)
}
