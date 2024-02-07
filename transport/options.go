package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	helper "github.com/antoha2/auth"
)

const authorizationHeader = "Authorization"

func (webImpl *webImpl) UserIdentify(ctx context.Context, r *http.Request) context.Context {

	header := r.Header.Get(authorizationHeader)
	if header == "" {
		newErr := "аутентификация - пустой заголовок"
		log.Println(newErr)
		return nil
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErr := "аутентификация - неправильный заголовок"
		log.Println(newErr)
		return nil
	}

	userId, err := webImpl.authService.ParseToken(headerParts[1])
	if err != nil {
		newErr := fmt.Sprintf("(auth) аутентификация - ошибка ParseToken()- %s", err)
		log.Println(newErr)
		return nil
	}

	if userId == 0 {
		newErr := "аутентификация - нет прав доступа / пользователь не найден"
		log.Println(newErr)
		return nil
	}

	userRoles := webImpl.authService.GetRoles(userId)
	if len(userRoles) == 0 {
		newErr := " не назначена роль"
		log.Println(newErr)
		return nil
	}

	ctx = context.WithValue(ctx, helper.USER_ID, userId)
	ctx = context.WithValue(ctx, helper.USER_ROLE, userRoles)
	return ctx

}
