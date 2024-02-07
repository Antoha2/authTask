package endpoints

import (
	"context"

	helper "github.com/antoha2/auth"
	authservice "github.com/antoha2/auth/service"
	"github.com/go-kit/kit/endpoint"
)

type SignUpAdminRequest struct {
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName  string `json:"lastname" gorm:"column:lastname"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password"`
}

type SignUpAdminResponse struct {
	UserId int `json:"user_id"`
}

const (
	roleAdmin = "admin"
	roleUser  = "user"
	roleDev   = "dev"
)

func MakeSignUpAdminEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(SignUpAdminRequest)

		inputUser := new(helper.User)
		inputRoles := new(helper.UsersRoles)

		inputUser.FirstName = req.FirstName
		inputUser.LastName = req.LastName
		inputUser.Password = req.Password
		inputUser.Username = req.Username

		inputRoles.Roles = append(inputRoles.Roles, roleAdmin)
		inputRoles.Roles = append(inputRoles.Roles, roleDev) // ?!? !!!!!!!!!!!!!!!!!!!!!!!!!!

		userId, err := s.CreateUser(inputUser, inputRoles)

		if err != nil {
			return nil, err
		}

		return SignUpAdminResponse{UserId: userId}, nil
	}
}
