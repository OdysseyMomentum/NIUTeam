package dbUtils

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBHelpers interface {
	CreateSession() (tableName *string, svc *dynamodb.DynamoDB)
	InsertItem(tableName *string, svc *dynamodb.DynamoDB, item interface{}) error
	CheckItemExists(tableName string, partitionKey string, sortKey string) (bool, error)
	GetItem(tableName string, partitionKey string, sortKey string) (interface{}, error)
}

type DynamoDBHelper struct{}

func (helper *DynamoDBHelper) CreateSession(inTableName string) (tableName *string, svc *dynamodb.DynamoDB) {
	tableName = aws.String(inTableName)
	sess := session.Must(session.NewSession())
	svc = dynamodb.New(sess)

	return tableName, svc
}

func (helper *DynamoDBHelper) InsertItem(tableName string, item map[string]*dynamodb.AttributeValue) error {
	cTableName, svc := helper.CreateSession(tableName)

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: cTableName,
	}
	_, err := svc.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (helper *DynamoDBHelper) CheckItemExists(tableName string, partitionKey string, sortKey string) (bool, error) {
	cTableName, svc := helper.CreateSession(tableName)
	fmt.Printf("The TableName: %v", tableName)
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		AttributesToGet: aws.StringSlice([]string{"PK"}),
		TableName:       cTableName,
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(partitionKey),
			},
			"SK": {
				S: aws.String(sortKey),
			},
		},
	})

	if err != nil {
		return false, err
	}

	if result.Item == nil {
		return false, nil
	}
	return true, nil
}

func (helper *DynamoDBHelper) GetItem(tableName string, partitionKey string, sortKey string) (interface{}, error) {
	cTableName, svc := helper.CreateSession(tableName)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: cTableName,
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(partitionKey),
			},
			"SK": {
				S: aws.String(sortKey),
			},
		},
	})

	if err != nil {
		return false, err
	}

	if result.Item == nil {
		return nil, errors.New("the item does not exist")
	}
	return result.Item, nil
}
