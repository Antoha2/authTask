package authrepository

import (
	helper "github.com/antoha2/auth"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *helper.User, userRoles *helper.UsersRoles) (int, error)
	UpdateUser(user *helper.User) error
	DeleteUser(userId int) error
	GetUser(username, password string) (*helper.User, error)
	GetRoles(id int) []string
}

type AuthPostgres struct {
	dbx *gorm.DB
}

func NewAuthPostgres(dbx *gorm.DB) *AuthPostgres {
	return &AuthPostgres{dbx: dbx}
}

/* type UserlistToRoles struct {
	user_id int
	role_id int
} */
