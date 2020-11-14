package users

import (
	"crypto/sha256"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/srbry/go-serverless-example/commons/addresses"
)

type User struct {
	UserId    string `json:"userId"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Birthdate uint64 `json:"birthdate"`
	addresses.Address
}

type UserObject struct {
	UserId      string `json:"userId"`
	Email       string `json:"email"`
	UOI         string `json:"uoi"`
	Role        string `json:"role"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

func GetUserPK(partitionKey interface{}) string {
	hash := sha256.New()
	hash.Write([]byte(partitionKey.(string)))

	emailHash := hash.Sum(nil)
	return fmt.Sprintf("USER#%x", emailHash)
}

func GetUserObjectSK(sortKey interface{}) string {
	return fmt.Sprintf("UO#%v", sortKey.(string))
}

func MarshallUserDetailsForDB(user User) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(struct {
		*User
		//OmitUserId string `json:"userId,omitempty"`

		PK string `json:"PK"`
		SK string `json:"SK"`
	}{
		User: &user,
		PK:   GetUserPK(user.Email),
		SK:   "DETAILS",
	})
}

func MarshallUserObjectsForDB(userObj UserObject) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(struct {
		*UserObject
		//OmitUserId string `json:"userId,omitempty"`

		PK string `json:"PK"`
		SK string `json:"SK"`
	}{
		UserObject: &userObj,
		PK:         GetUserPK(userObj.Email),
		SK:         GetUserObjectSK(userObj.UOI),
	})
}
