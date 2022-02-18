package common

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/voto-vote/common/db"
)

type Links struct {
	Self string `json:"self"`
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type MetaData struct {
	Count int   `json:"count"`
	Total int   `json:"total"`
	Page  int   `json:"page"`
	Links Links `json:"_links"`
}

type HandlerParams struct {
	DbAccessor db.PostgresConnector
	Ctx        context.Context
	Request    events.APIGatewayProxyRequest
}
