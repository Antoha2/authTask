package endpoints

import (
	"context"
	"errors"
	"log"

	helper "github.com/antoha2/auth"
	authservice "github.com/antoha2/auth/service"
	"github.com/go-kit/kit/endpoint"
)

type DeleteUserRequest struct {
}

type DeleteUserResponse struct {
	UserId int `json:"user_id"`
}

func MakeDeleteUserEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		var userId int
		var ok bool
		if userId, ok = ctx.Value(helper.USER_ID).(int); !ok {
			newErr := "UserId не найден"
			log.Println(newErr)
			return nil, errors.New(newErr)
		}
		if err := s.DeleteUser(userId); err != nil {
			return nil, err
		}
		return DeleteUserResponse{UserId: userId}, nil
	}
}
