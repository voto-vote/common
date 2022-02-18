package auth

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type UserCustomClaims struct {
	IsAdmin     bool  `json:"isAdmin"`
	Creator     []int `json:"creator"`
	Trustperson []int `json:"trustperson"`
	Candidate   []int `json:"candidate"`
	Media       []int `json:"media"`
	VOTO_id     int64 `json:"VOTO_id"`
}

type Firebase struct {
	Auth *auth.Client
	App  *firebase.App
	Ctx  context.Context
}
