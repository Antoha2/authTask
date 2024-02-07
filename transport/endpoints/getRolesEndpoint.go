package endpoints

import (
	"context"

	authservice "github.com/antoha2/auth/service"
	"github.com/go-kit/kit/endpoint"
)

type GetRolesRequest struct {
	UserId int `json:"user_id"`
}

type GetRolesResponse struct {
	Roles []string `json:"roles"`
}

func MakeGetRolesEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetRolesRequest)
		roles := s.GetRoles(req.UserId)

		return GetRolesResponse{roles}, nil
	}
}
