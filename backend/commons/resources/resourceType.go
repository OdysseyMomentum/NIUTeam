package resources

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	ResourceKeyPrefix = "RSC"
)

type Resource struct {
	ResourceId   string `json:"resourceId"`
	DisplayName  string `json:"displayName"`
	Description  string `json:"description"`
	CreatedAt    int64  `json:"createdAt"`
	ResourceType string `json:"resourceType"`
	Vendor       string `json:"vendor"`
	Meta         interface{} `json:"meta"`
}

type DocumentMeta struct {
	Filename string `json:"filename"`
	Filesize uint64 `json:"filesize"`
	Filetype string `json:"filetype"`
	Access string `json:"access"`
}

func GetResourcePK(partitionKey interface{}) string {
	return fmt.Sprintf("RSC#%v", partitionKey.(string))
}

func MarshallResourceDetailsForDB(resource Resource) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(struct {
		*Resource
		PK string `json:"PK"`
		SK string `json:"SK"`
	}{
		Resource: &resource,
		PK:       GetResourcePK(resource.ResourceId),
		SK:       "DETAILS",
	})
}
