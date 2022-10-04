package encryption

import (
	"github.com/albertopformoso/inventory/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

func SignedLoginToken(u *model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})

	return token.SignedString([]byte(key))
}
