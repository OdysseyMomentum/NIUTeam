package main

import (
	"github.com/aws/aws-lambda-go/events"
	"os"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {

	os.Setenv("AUTH0_JWKS_URI", "https://filedgr-dev.eu.auth0.com/.well-known/jwks.json")
	os.Setenv("AUTH0_AUDIENCE", "https://dev.streams.filedgr.com")

	type args struct {
		request events.APIGatewayCustomAuthorizerRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayCustomAuthorizerResponse
		wantErr bool
	}{
		{
			name: "Testing wrong bearer prefix.",
			args: args{request: events.APIGatewayCustomAuthorizerRequest{
				Type:               "",
				AuthorizationToken: "Dummy Token",
				MethodArn:          "",
			}},
			wantErr: true,
		},
		{
			name: "Correct Bearer.",
			args: args{request: events.APIGatewayCustomAuthorizerRequest{
				Type:               "",
				AuthorizationToken: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlVTdGxkTzFoOW5wZ0p6OTFwdGNELSJ9." +
					"eyJpc3MiOiJodHRwczovL2ZpbGVkZ3ItZGV2LmV1LmF1dGgwLmNvbS8iLCJzdWIiOiJhdXRoMHw1ZjkzM2VkZWJiZGE1MD" +
					"AwNmExMWVmZGEiLCJhdWQiOiJodHRwczovL2Rldi5zdHJlYW1zLmZpbGVkZ3IuY29tIiwiaWF0IjoxNjAzNDg2NzQ0LCJle" +
					"HAiOjE2MDM1NzMxNDQsImF6cCI6Im5aY0RIRGE3Mkc0NEp6a2FaRnRURFkwcEZUYzFnRzhNIiwiZ3R5IjoicGFzc3dvcmQ" +
					"ifQ.aVO538lHqET8Hb_iCU0az7pDshWlWK2IdEvsRou5geCLUzVUFC5iICIZ00gu9TpgsEoywWdUrpifpRjwsi58pJBkIhvD" +
					"hysZdosv5y_APIGfebedgwOiGtzo2U5sb2k45EZ02UHQcQNpS_n6ktQPThbTk32gLF7qBXmzVr5JZAKoQntCOmRFgjXE9oo" +
					"wow-7_-u4lKpzdgrK6zgwghuososq1A05IEMJtkld5KCzg5F_fxz_TbhYEvQRYDNj-l20To6-qH5sGU-ZeeDbfeGJKDBUdAf" +
					"k6ujGoOQb7dpKmpQSQPuNPfJlJQ5mYE19GFYVqd4DipqWkN7ZC7Gkm5LRMg",
				MethodArn:          "",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePolicy(t *testing.T) {
	type args struct {
		principalID string
		effect      string
		resource    string
	}
	tests := []struct {
		name string
		args args
		want events.APIGatewayCustomAuthorizerResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePolicy(tt.args.principalID, tt.args.effect, tt.args.resource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}