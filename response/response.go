package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/voto-vote/common"
)

var CORS_HEADERS = map[string]string{
	"Access-Control-Allow-Headers":     "Content-Type,X-Amz-Date,X-Amz-Security-Token,Authorization,X-Api-Key,X-Requested-With,Accept,Access-Control-Allow-Methods,Access-Control-Allow-Origin,Access-Control-Allow-Headers",
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Methods":     "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT",
	"Access-Control-Allow-Credentials": "true",
	"X-Requested-With":                 "*",
}

func HandleError(status int32, errorMsg string) (events.APIGatewayProxyResponse, error) {
	errResponse := common.ErrorResponse{
		Status:  status,
		Message: errorMsg,
	}
	jsoned, err := json.Marshal(errResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    CORS_HEADERS,
			StatusCode: int(status),
			Body:       string(jsoned),
		}, err
	}
	return events.APIGatewayProxyResponse{
		Headers:    CORS_HEADERS,
		StatusCode: int(status),
		Body:       string(jsoned),
	}, nil
}

func HandleSuccess(code int, body string) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       body,
		Headers:    CORS_HEADERS,
	}, nil
}
