package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/users"
	"os"
)

type objectLister struct{}

var objectList objectLister

func init() {
	objectList = objectLister{}
}

func handleObjectList(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	user := users.User{}
	user.Email, user.UserId = commons.GetUserProfileInfo(request)

	results, err := dbUtils.ListObjectsForUser(os.Getenv("TABLE_NAME"), users.GetUserPK(user.Email))
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	resp, err := commons.BuildResponse(200, results)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	return resp, nil
}

func main() {
	lambda.Start(handleObjectList)
}
