package dbUtils

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/srbry/go-serverless-example/commons/users"
)

func GetUserById(tableName string, userId string) (*users.User, error) {
	dbHelper := DynamoDBHelper{}
	cTableName, svc := dbHelper.CreateSession(tableName)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: cTableName,
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(userId),
			},
			"SK": {
				S: aws.String("DETAILS"),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return &users.User{}, errors.New("the item was not found")
	}

	resultUser := users.User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &resultUser)
	if err != nil {
		return nil, err
	}
	return &resultUser, nil
}
