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
	"os"
	"time"
)

type resourceInserter struct{}
type userObjectGetter struct{}
type objectResourceInserter struct{}

var resourceInsert resourceInserter
var userObjectGet userObjectGetter
var objectResourceInsert objectResourceInserter

func init() {
	resourceInsert = resourceInserter{}
	userObjectGet = userObjectGetter{}
	objectResourceInsert = objectResourceInserter{}
}

func main() {
	lambda.Start(handleAddResource)
}

func handleAddResource(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := users.User{}
	user.Email, user.UserId = commons.GetUserProfileInfo(request)

	fmt.Printf("The body: %v\n", request.Body)

	inputStruct := struct {
		Object   objects.Object     `json:"object"`
		Resource resources.Resource `json:"resource"`
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

	fmt.Printf("The userObj: %+v\n", userObj)

	//Insert the resource
	resource := inputStruct.Resource
	resource.CreatedAt = time.Now().Unix()
	err = resourceInsert.InsertItem(resource)
	if err != nil {
		return commons.BuildResponse(200, err.Error())
	}

	//Insert the Object Resource
	objectResource := objects.ObjectResource{
		UOI:          object.UOI,
		ResourceId:   resource.ResourceId,
		DisplayName:  resource.DisplayName,
		CreatedAt:    resource.CreatedAt,
		ResourceType: resource.ResourceType,
	}
	err = objectResourceInsert.InsertItem(objectResource)
	if err != nil {
		return commons.BuildResponse(500, err)
	}

	//TODO: Get all the users connected to this obj
	//TODO: Batch insert all the user Resources

	return commons.BuildResponse(200, resource)
}

func (resourceInsert *resourceInserter) InsertItem(item interface{}) error {
	resource := item.(resources.Resource)

	av, err := resources.MarshallResourceDetailsForDB(resource)
	if err != nil {
		return err
	}

	dbHelper := dbUtils.DynamoDBHelper{}
	dbHelper.InsertItem(os.Getenv("TABLE_NAME"), av)

	return nil
}

func (objectResourceInsert *objectResourceInserter) InsertItem(item interface{}) error {
	resource := item.(objects.ObjectResource)

	av, err := objects.MarshallObjectResourcesForDB(resource)
	if err != nil {
		return err
	}

	dbHelper := dbUtils.DynamoDBHelper{}
	dbHelper.InsertItem(os.Getenv("TABLE_NAME"), av)

	return nil
}

func (userObjectExistsCheck *userObjectGetter) GetItem(partitionKey interface{}, sortKey interface{}) (interface{}, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	return dbHelper.CheckItemExists(os.Getenv("TABLE_NAME"),
		partitionKey.(string), sortKey.(string))
}
