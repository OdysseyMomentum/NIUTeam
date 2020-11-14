package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	//Getting the authToken
	authToken := request.AuthorizationToken
	tokenSlice := strings.Split(authToken, " ")

	//Check the format of the authToken
	var bearerToken string
	if len(tokenSlice) > 1 {
		bearerToken = tokenSlice[0]
	}
	if strings.ToLower(bearerToken) != "bearer" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	//Verify the authToken
	token, err := verifyToken(tokenSlice[1])
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}
	if !token.Valid {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}
	claims := token.Claims.(jwt.MapClaims)
	return generatePolicy(claims, "Allow", request.MethodArn), nil
}

func main() {
	lambda.Start(handler)
}

func generatePolicy(claims jwt.MapClaims, effect, resource string) events.APIGatewayCustomAuthorizerResponse {

	mailClaim := fmt.Sprintf("%v/email", os.Getenv("AUTH0_TOKEN_NAMESPACE"))

	principalID := claims["sub"].(string)
	email := claims[mailClaim].(string)

	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
		authResponse.Context = map[string]interface{}{
			"email":  email,
			"userId": principalID,
		}
	}
	return authResponse
}

func generateDeny() events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{}
	authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
		Version: "2012-10-17",
		Statement: []events.IAMPolicyStatement{
			{
				Action:   []string{"execute-api:Invoke"},
				Effect:   "Deny",
				Resource: []string{"*"},
			},
		},
	}
	return authResponse
}

func getJwks(token *jwt.Token) (interface{}, error) {
	//Get the KeyId
	keyID, in := token.Header["kid"].(string)
	if !in {
		return nil, errors.New("the token header should contain a keyId")
	}

	jwksUri := os.Getenv("AUTH0_JWKS_URI")
	set, err := jwk.FetchHTTP(jwksUri)
	if err != nil {
		return nil, err
	}

	var key interface{}
	if keys := set.LookupKeyID(keyID); len(keys) == 1 {
		err = keys[0].Raw(&key)
		if err != nil {
			return nil, err
		}
		return key, nil
	}
	return nil, fmt.Errorf("unable to find key %q", keyID)
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, getJwks)
	if err != nil {
		return nil, err
	}
	return token, nil
}
