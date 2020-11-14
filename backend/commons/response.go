package commons

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

var typesMap = map[string]string{
	".pdf":  "application/pdf",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".jpeg":  "image/jpeg",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".xls":  "application/vnd.ms-excel",
	".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".ppt":  " application/vnd.ms-powerpoint",
	".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
}

func BuildResponse(responseCode int, jsonBody interface{}) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(jsonBody)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse{
		StatusCode:      responseCode,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}

func GetContentTypeForFile(extension string) string {
	contentType, ok := typesMap[extension]
	if !ok {
		return "text/plain"
	}
	return contentType
}
