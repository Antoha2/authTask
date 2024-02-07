package authservice

import (
	"time"

	helper "github.com/antoha2/auth"
	authRepository "github.com/antoha2/auth/service/authRepository"
)

const (
	salt       = "aW1;"
	signingKey = "Bgt5"
	tokenTTL   = 12 * time.Hour
)

type Authorization interface {
	CreateUser(user *helper.User, userRoles *helper.UsersRoles) (int, error)
	UpdateUser(user *helper.User) error
	DeleteUser(userId int) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetRoles(id int) []string
}

type AuthService struct {
	authRep authRepository.Authorization
}

func NewAuthService(authRep authRepository.Authorization) *AuthService {
	return &AuthService{
		authRep: authRep,
	}
}
