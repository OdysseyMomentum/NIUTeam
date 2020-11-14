package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/users"
	"os"
)


type userObjectGetter struct{}
type userObjectInserter struct {}

var userObjectGet userObjectGetter
var userObjectInsert userObjectInserter

func init() {
	userObjectGet = userObjectGetter{}
	userObjectInsert = userObjectInserter{}
}


func handleUserObjectInsert(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := users.User{}
	user.Email, user.UserId = commons.GetUserProfileInfo(request)
	fmt.Printf("The body: %v\n", request.Body)

	inputStruct := struct {
		Object   objects.Object     `json:"object"`
		User users.User `json:"user"`
	}{}
	err := commons.UnmarshallRequestBody(request.Body, &inputStruct)
	if err != nil {
		return commons.BuildResponse(200, err.Error())
	}

	fmt.Printf("The inputStruct: %+v\n", inputStruct)

	object := inputStruct.Object

	//Check if user is allowed to
	userObj, err := userObjectGet.GetItem(users.GetUserPK(user.Email), users.GetUserObjectSK(object.UOI))
	if err != nil {
		return commons.BuildResponse(200, err.Error())
	}
	currentUserObj := users.UserObject{}
	dynamodbattribute.UnmarshalMap(userObj.(map[string]*dynamodb.AttributeValue), &currentUserObj)

	fmt.Printf("The userObj: %+v\n", userObj)

	userObjToInsert := users.UserObject{
		UserId:      users.GetUserPK(inputStruct.User.Email),
		Email:       inputStruct.User.Email,
		UOI:         inputStruct.Object.UOI,
		Role:        "Admin",
		DisplayName: currentUserObj.DisplayName,
		Description: currentUserObj.Description,
	}

	userObjectInsert.InsertItem(userObjToInsert)

	return commons.BuildResponse(200, userObjToInsert)
}

func (userObjectInsert *userObjectInserter) InsertItem(item interface{}) error {
	userObj := item.(users.UserObject)
	fmt.Printf("The user/userObj: %+v", userObj)
	av, err := users.MarshallUserObjectsForDB(userObj)
	if err != nil {
		return err
	}

	dbHelper := dbUtils.DynamoDBHelper{}
	err = dbHelper.InsertItem(os.Getenv("TABLE_NAME"), av)
	return err
}


func (userObjectExistsCheck *userObjectGetter) GetItem(partitionKey interface{}, sortKey interface{}) (interface{}, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	return dbHelper.GetItem(os.Getenv("TABLE_NAME"),
		partitionKey.(string), sortKey.(string))
}


func main() {
	lambda.Start(handleUserObjectInsert)
}
