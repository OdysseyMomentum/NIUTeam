package objects

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/srbry/go-serverless-example/commons/addresses"
	"github.com/srbry/go-serverless-example/commons/resources"
)

type Object struct {
	UOI            string                   `json:"uoi"`
	DisplayName    string                   `json:"displayName"`
	Description    string                   `json:"description"`
	GeoCoordinates addresses.GeoCoordinates `json:"geoCoordinates"`
	addresses.Address
}

type ObjectResource struct {
	UOI          string `json:"uoi"`
	ResourceId   string `json:"resourceId"`
	DisplayName  string `json:"displayName"`
	CreatedAt    int64  `json:"createdAt"`
	ResourceType string `json:"resourceType"`
	Meta interface{} `json:"meta"`
}

func GetObjectPK(partitionKey interface{}) string {
	return fmt.Sprintf("UO#%v", partitionKey.(string))
}

func GetObjectResourceSK(sortKey interface{}) string {
	return fmt.Sprintf("%v#%v", resources.ResourceKeyPrefix, sortKey.(string))
}

func MarshallObjectDetailsForDB(object Object) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(struct {
		*Object
		PK string `json:"PK"`
		SK string `json:"SK"`
	}{
		Object: &object,
		PK:     GetObjectPK(object.UOI),
		SK:     "DETAILS",
	})
}

func MarshallObjectResourcesForDB(objectRsc ObjectResource) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(struct {
		*ObjectResource
		PK string `json:"PK"`
		SK string `json:"SK"`
	}{
		ObjectResource: &objectRsc,
		PK:             GetObjectPK(objectRsc.UOI),
		SK:             GetObjectResourceSK(objectRsc.ResourceId),
	})
}
