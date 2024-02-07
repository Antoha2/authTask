package endpoints

import (
	"context"

	authservice "github.com/antoha2/auth/service"
	"github.com/go-kit/kit/endpoint"
)

type ParseTokenRequest struct {
	Token string `json:"token"`
}

type ParseTokenResponse struct {
	UserId int `json:"user_id"`
}

func MakeParseTokenEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(ParseTokenRequest)

		userId, err := s.ParseToken(req.Token)
		if err != nil {
			return nil, err
		}

		return ParseTokenResponse{userId}, nil
	}
}
