package dbUtils

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/sensordata"
	"strconv"
)

func ListResourcesForObject(tableName string, PK string, skPrefix string) (interface{}, error) {
	dbHelper := DynamoDBHelper{}
	cTableName, svc := dbHelper.CreateSession(tableName)

	queryInput := &dynamodb.QueryInput{
		TableName: cTableName,
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partitionKey": {
				S: aws.String(PK),
			},
			":sortKey": {
				S: aws.String(skPrefix),
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

	var userObjs []objects.ObjectResource
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &userObjs)
	if err != nil {
		return nil, err
	}
	return &userObjs, nil
}

func GetSensorData(tableName string, PK string, begin uint64, end uint64) (interface{}, error) {
	dbHelper := DynamoDBHelper{}
	cTableName, svc := dbHelper.CreateSession(tableName)
	fmt.Printf("The Table Name for IoT: %v\n", tableName)
	queryInput := &dynamodb.QueryInput{
		TableName: cTableName,
		ProjectionExpression: aws.String("device_id, #timestamp, altitude, co2_conc, humidity, illuminance,	motion_count, pressure, sound_level, temperature, voc_conc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partitionKey": {
				S: aws.String(PK),
			},
			":begin": {
				N: aws.String(strconv.FormatUint(begin, 10)),
			},
			":end": {
				N: aws.String(strconv.FormatUint(end, 10)),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#timestamp": aws.String("timestamp"),
		},
		KeyConditionExpression: aws.String("device_id = :partitionKey AND #timestamp BETWEEN :begin AND :end"),
	}

	result, err := svc.Query(queryInput)
	if err != nil {
		return nil, err
	}

	if result.Items == nil {
		return nil, errors.New("no items found")
	}

	var iotData []sensordata.CooperDataType
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &iotData)
	if err != nil {
		return nil, err
	}
	return &iotData, nil
}
