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

type userGetter struct{}
type userInserter struct{}

var userGet userGetter
var userInsert userInserter

func init() {
	userGet = userGetter{}
	userInsert = userInserter{}
}

func (userGet *userGetter) GetItem(partitionKey interface{}, rangeKey interface{}) (interface{}, error) {
	user, err := dbUtils.GetUserById(os.Getenv("TABLE_NAME"), users.GetUserPK(partitionKey))
	return user, err
}

func (userInsert *userInserter) InsertItem(item interface{}) error {
	user := item.(*users.User)
	av, err := users.MarshallUserDetailsForDB(*user)
	if err != nil {
		return err
	}

	dbHelper := dbUtils.DynamoDBHelper{}
	err = dbHelper.InsertItem(os.Getenv("TABLE_NAME"), av)
	return err
}

func userGetHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	user := &users.User{}
	user.Email, user.UserId = commons.GetUserProfileInfo(request)

	getResult, err := userGet.GetItem(user.Email, nil)
	if err != nil {
		if getResult != nil {
			err = userInsert.InsertItem(user)
			if err != nil {
				return events.APIGatewayProxyResponse{StatusCode: 500}, err
			}
		} else {
			return events.APIGatewayProxyResponse{StatusCode: 500}, err
		}
	} else {
		user = getResult.(*users.User)
	}

	response, err := commons.BuildResponse(200, user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	return response, nil
}

func main() {
	lambda.Start(userGetHandler)
}
