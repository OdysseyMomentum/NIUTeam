package commons

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func UnmarshallRequestBody(body string, data interface{}) error {
	return json.Unmarshal([]byte(body), data)
}

func GetUserProfileInfo(request events.APIGatewayProxyRequest) (email string, userId string) {
	return request.RequestContext.Authorizer["email"].(string), request.RequestContext.Authorizer["userId"].(string)
}
