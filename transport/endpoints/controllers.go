package endpoints

import (
	authservice "github.com/antoha2/auth/service"

	"github.com/go-kit/kit/endpoint"
)

type AuthEndpoints struct {
	SignIn      endpoint.Endpoint
	SignUpAdmin endpoint.Endpoint
	SignUpUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	ParseToken  endpoint.Endpoint
	GetRoles    endpoint.Endpoint
}

func MakeAuthEndpoints(s authservice.Authorization) *AuthEndpoints {
	return &AuthEndpoints{
		SignIn:      MakeSignInEndpoint(s),
		SignUpAdmin: MakeSignUpAdminEndpoint(s),
		SignUpUser:  MakeSignUpUserEndpoint(s),
		DeleteUser:  MakeDeleteUserEndpoint(s),
		UpdateUser:  MakeUpdateUserEndpoint(s),
		ParseToken:  MakeParseTokenEndpoint(s),
		GetRoles:    MakeGetRolesEndpoint(s),
	}
}
