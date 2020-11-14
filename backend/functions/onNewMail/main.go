package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/DusanKasan/parsemail"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/google/uuid"
	"github.com/srbry/go-serverless-example/commons/dbUtils"
	"github.com/srbry/go-serverless-example/commons/objects"
	"github.com/srbry/go-serverless-example/commons/resources"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type objectExistsChecker struct{}
type resourceInserter struct{}
type objectResourceInserter struct{}

var objectExistsCheck objectExistsChecker
var resourceInsert resourceInserter
var objectResourceInsert objectResourceInserter

func init() {
	objectExistsCheck = objectExistsChecker{}
	resourceInsert = resourceInserter{}
	objectResourceInsert = objectResourceInserter{}
}

func (objectExistsCheck *objectExistsChecker) CheckItemExists(partitionKey interface{}, sortKey interface{}) (bool, error) {
	dbHelper := dbUtils.DynamoDBHelper{}
	return dbHelper.CheckItemExists(os.Getenv("TABLE_NAME"), partitionKey.(string), sortKey.(string))
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

func handleOnNewMail(ctx context.Context, bucketEvent events.S3Event) error {

	svc := s3iface.S3API(s3.New(session.Must(session.NewSession())))
	for _, v := range bucketEvent.Records {
		key := v.S3.Object.Key

		obj, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(v.S3.Bucket.Name),
			Key:    aws.String(key),
		})
		if err != nil {
			return fmt.Errorf("error in downloading %s from S3: %s\n", key, err)
		}

		body, err := ioutil.ReadAll(obj.Body)
		if err != nil {
			return fmt.Errorf("error in reading file %s: %s\n", key, err)
		}

		email, err := parsemail.Parse(bytes.NewReader(body))
		if err != nil {
			return fmt.Errorf("error in reading file %s: %s\n", key, err)
		}

		//Get the ID part from the Email:
		to := email.To
		for _, mail := range to {
			tokens := strings.Split(mail.Address, "@")
			if tokens[1] == os.Getenv("APP_DOMAIN") {
				//We got the mail for the UOI, we now check does this UOI exist?
				uoi := tokens[0]
				exists, err := objectExistsCheck.CheckItemExists(objects.GetObjectPK(tokens[0]), "DETAILS")
				if err != nil {
					return err
				}
				if !exists {
					return errors.New("this object does not exist")
				}

				//If it exists we create the resource.
				fmt.Printf("The object exists: %v\n", tokens[0])

				fmt.Printf("Embedded files: %v\n", email.EmbeddedFiles)

				for _, a := range email.Attachments {
					fmt.Printf("Filename: %v\n", a.Filename)
					inputBuf := make([]byte, 20480)
					var outputBuf bytes.Buffer
					outputWriter := bufio.NewWriter(&outputBuf)
					var totalBytes uint64 = 0

					for {
						read, err := a.Data.Read(inputBuf)
						if read > 0 {
							totalBytes += uint64(read)
							outputWriter.Write(inputBuf[:read])
						}
						if err != nil {
							if err != io.EOF {
								return err
							}
							break
						}
					}

					//We create the resource
					fileTokens := strings.Split(a.Filename, ".")
					resource := resources.Resource{
						ResourceId:   uuid.New().String(),
						DisplayName:  a.Filename,
						CreatedAt:    email.Date.Unix(),
						ResourceType: "DOCUMENT",
						Meta: resources.DocumentMeta{
							Filename: a.Filename,
							Filesize: totalBytes,
							Filetype: "." + fileTokens[len(fileTokens)-1],
						},
					}
					err := resourceInsert.InsertItem(resource)
					if err != nil {
						return err
					}

					objectResource := objects.ObjectResource{
						UOI:          uoi,
						ResourceId:   resource.ResourceId,
						DisplayName:  resource.DisplayName,
						CreatedAt:    resource.CreatedAt,
						ResourceType: "DOCUMENT",
						Meta: resources.DocumentMeta{
							Filename: a.Filename,
							Filesize: totalBytes,
							Filetype: "." + fileTokens[len(fileTokens)-1],
						},
					}
					err = objectResourceInsert.InsertItem(objectResource)
					if err != nil {
						return err
					}

					//Storing to S3
					putOutput, err := svc.PutObject(&s3.PutObjectInput{
						Bucket: aws.String(os.Getenv("BUCKET_NAME")),
						Key:    aws.String(fmt.Sprintf("%v/%v", tokens[0], resource.ResourceId)),
						Body:   bytes.NewReader(outputBuf.Bytes()),
					})
					if err != nil {
						fmt.Errorf("error in reading file %s: %s\n", key, err)
						return nil
					}

					fmt.Errorf("Put output: %v\n", putOutput)
				}
			}
		}
	}
	return nil
}

func main() {
	lambda.Start(handleOnNewMail)
}
