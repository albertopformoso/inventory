package encryption

import (
	"time"

	"github.com/albertopformoso/inventory/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

func SignedLoginToken(u *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString([]byte(key))
}

func ParseLoginJWT(value string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
