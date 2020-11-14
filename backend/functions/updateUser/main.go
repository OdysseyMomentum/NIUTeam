package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/srbry/go-serverless-example/commons"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/users"
	"os"
	"strconv"
)

type userUpdater struct{}

var userUpdate userUpdater

func init() {
	userUpdate = userUpdater{}
}

func (userUpdater *userUpdater) UpdateItem(item interface{}) error {
	dbHelper := dbUtils.DynamoDBHelper{}
	tableName, svc := dbHelper.CreateSession(os.Getenv("TABLE_NAME"))

	user := item.(users.User)

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":firstname": {
				S: aws.String(user.Firstname),
			},
			":lastname": {
				S: aws.String(user.Lastname),
			},
			":birthdate": {
				N: aws.String(strconv.FormatUint(user.Birthdate, 10)),
			},
			":streetName": {
				S: aws.String(user.StreetName),
			},
			":streetNumber": {
				S: aws.String(user.StreetNumber),
			},
			":zipcode": {
				S: aws.String(user.Zipcode),
			},
			":city": {
				S: aws.String(user.City),
			},
			":country": {
				S: aws.String(user.Country),
			},
		},
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(users.GetUserPK(user.Email)),
			},
			"SK": {
				S: aws.String("DETAILS"),
			},
		},
		ReturnValues: aws.String("NONE"),
		UpdateExpression: aws.String("set firstname = :firstname, " +
			"lastname = :lastname, " +
			"birthdate = :birthdate, " +
			"streetName = :streetName, " +
			"streetNumber = :streetNumber, " +
			"zipcode = :zipcode, " +
			"city = :city, " +
			"country = :country"),
	}

	output, err := svc.UpdateItem(input)
	fmt.Printf("The output: %v", output)
	if err != nil {
		return err
	}

	return nil
}

func UserUpdateHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := users.User{}
	commons.UnmarshallRequestBody(request.Body, &user)
	user.Email, user.UserId = commons.GetUserProfileInfo(request)

	err := userUpdate.UpdateItem(user)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, err := commons.BuildResponse(200, user)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.New("an error occurred when building the response")
	}
	return response, nil
}

func main() {
	lambda.Start(UserUpdateHandler)
}
