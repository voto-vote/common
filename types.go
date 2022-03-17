package common

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/voto-vote/common/db"
)

type Links struct {
	Self string `json:"self"`
	Prev string `json:"prev_page_url"`
	Next string `json:"next_page_url"`
}

type Pagination struct {
	Total       int   `json:"total"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
	From        int   `json:"from"`
	To          int   `json:"to"`
	Links       Links `json:"_links"`
}

type HandlerParams struct {
	DbAccessor db.PostgresConnector
	Ctx        context.Context
	Request    events.APIGatewayProxyRequest
}

type ErrorResponse struct {
	Status  int32
	Message string
}

type VOTODate time.Time
