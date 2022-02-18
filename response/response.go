package response

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/voto-vote/common"
)

func HandleError(status int32, errorMsg string) (events.APIGatewayProxyResponse, error) {
	errResponse := common.ErrorResponse{
		Status:  status,
		Message: errorMsg,
	}
	jsoned, _ := json.Marshal(errResponse)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       string(jsoned),
	}, nil
}

func HandleSuccess(code int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       body,
	}, nil
}
