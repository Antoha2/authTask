package endpoints

import (
	"context"

	helper "github.com/antoha2/auth"
	authservice "github.com/antoha2/auth/service"
	"github.com/go-kit/kit/endpoint"
)

type SignUpUserRequest struct {
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName  string `json:"lastname" gorm:"column:lastname"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password"`
}

type SignUpUserResponse struct {
	UserId int `json:"user_id"`
}

func MakeSignUpUserEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(SignUpUserRequest)

		inputUser := new(helper.User)
		inputRoles := new(helper.UsersRoles)

		inputUser.FirstName = req.FirstName
		inputUser.LastName = req.LastName
		inputUser.Password = req.Password
		inputUser.Username = req.Username

		inputRoles.Roles = append(inputRoles.Roles, roleUser)

		userId, err := s.CreateUser(inputUser, inputRoles)
		if err != nil {
			return nil, err
		}
		return SignUpUserResponse{UserId: userId}, nil
	}
}
