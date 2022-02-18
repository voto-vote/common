package response

import "github.com/aws/aws-lambda-go/events"

func HandleSuccess(code int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       body,
	}, nil
}
