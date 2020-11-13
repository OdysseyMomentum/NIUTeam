package dbUtils

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/users"
)

func GetObjectById(tableName string, uoi string) (*objects.Object, error) {
	dbHelper := DynamoDBHelper{}
	cTableName, svc := dbHelper.CreateSession(tableName)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: cTableName,
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(uoi),
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
		return nil, errors.New("the item was not found")
	}

	resultObject := objects.Object{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &resultObject)
	if err != nil {
		return nil, err
	}
	return &resultObject, nil
}

func ListObjectsForUser(tableName string, userKey string) (interface{}, error) {
	dbHelper := DynamoDBHelper{}
	cTableName, svc := dbHelper.CreateSession(tableName)

	queryInput := &dynamodb.QueryInput{
		TableName: cTableName,
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partitionKey": {
				S: aws.String(userKey),
			},
			":sortKey": {
				S: aws.String("UO#"),
			},
		},
		KeyConditionExpression: aws.String("PK = :partitionKey AND begins_with(SK, :sortKey)"),
	}

	result, err := svc.Query(queryInput)
	if err != nil {
		return nil, err
	}

	if result.Items == nil {
		return nil, errors.New("no items found")
	}

	userObjs := []users.UserObject{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &userObjs)
	if err != nil {
		return nil, err
	}
	return &userObjs, nil
}
