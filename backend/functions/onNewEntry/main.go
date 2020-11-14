package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	sns2 "github.com/aws/aws-sdk-go/service/sns"
)

func handleNewEntry(ctx context.Context, event events.DynamoDBEvent) (error) {
	//fmt.Printf("The records: %v\n", event )
	sess := session.Must(session.NewSession())
	sns := sns2.New(sess)
	_ = sns

	for _, record := range event.Records {
		//var myStruct map[string]interface{}
		//err := dynamodbattribute.Unmarshal(record.Change.NewImage, &myStruct)
		//if err != nil {
		//	return err
		//}

		fmt.Printf("The REcord: %v\n", record.Change.NewImage)

	}


	//snsInput := &sns2.PublishInput{
	//	Message:                aws.String("The CO2 level for is to high! Please react."),
	//	TopicArn:               aws.String(os.Getenv("SNS_TOPIC_ALERT")),
	//
	//}
	//_, err := sns.Publish(snsInput)
	//
	//return err
	return nil
}

func main() {
	lambda.Start(handleNewEntry)
}
