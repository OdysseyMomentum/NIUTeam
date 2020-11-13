package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/users"
	"os"
)

type objectExistsChecker struct{}
type objectInserter struct{}
type userObjectInserter struct{}

var objectExistsCheck objectExistsChecker
var objectInsert objectInserter
var userObjectInsert userObjectInserter

func init() {
	objectExistsCheck = objectExistsChecker{}
	objectInsert = objectInserter{}
	userObjectInsert = userObjectInserter{}
}

func (objectGet *objectExistsChecker) CheckItemExists(partitionKey interface{},
	sortKey interface{}) (bool, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	return dbHelper.CheckItemExists(os.Getenv("TABLE_NAME"),
		partitionKey.(string), sortKey.(string))
}

func (objectInsert *objectInserter) InsertItem(item interface{}) error {

	object := item.(objects.Object)
	fmt.Printf("The object: %+v", object)
	av, err := objects.MarshallObjectDetailsForDB(object)
	if err != nil {
		return err
	}

	dbHelper := dbUtils.DynamoDBHelper{}
	err = dbHelper.InsertItem(os.Getenv("TABLE_NAME"), av)
	return err
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

func handleAddObject(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	user := users.User{}
	user.Email, user.UserId = commons.GetUserProfileInfo(request)

	uo := objects.Object{}
	commons.UnmarshallRequestBody(request.Body, &uo)

	exists, err := objectExistsCheck.CheckItemExists(objects.GetObjectPK(uo.UOI), "DETAILS")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	if exists {
		return events.APIGatewayProxyResponse{StatusCode: 500}, errors.New(fmt.Sprintf("The object with UOI: %v already exists", uo.UOI))
	}

	err = objectInsert.InsertItem(uo)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	userObj := users.UserObject{
		UserId:      user.UserId,
		Email:       user.Email,
		UOI:         uo.UOI,
		Role:        "OWNER",
		DisplayName: uo.DisplayName,
		Description: uo.Description,
	}
	err = userObjectInsert.InsertItem(userObj)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	fmt.Printf("The userObj: %v", userObj)

	resp, err := commons.BuildResponse(200, uo)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	return resp, nil
}

func main() {
	lambda.Start(handleAddObject)
}
