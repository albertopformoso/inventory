package middleware

import (
	"net/http"

	"github.com/albertopformoso/inventory/encryption"
	"github.com/albertopformoso/inventory/internal/helper"
	"github.com/albertopformoso/inventory/logger"
	"github.com/labstack/echo/v4"
)

var log = logger.New()

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get auth token from cookie
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			log.Err(err).Msg("Retrieve token faild")
			res := helper.BuildErrorResopnse("retrive token failed", "Unauthorized", helper.EmptyObj{})
			return c.JSONPretty(http.StatusUnauthorized, res, "  ")
		}

		// Parse the JWT token
		claims, err := encryption.ParseLoginJWT(cookie.Value)
		if err != nil {
			log.Err(err).Msg("Unauthorized")
			if err.Error() == "Token is expired" {
				res := helper.BuildErrorResopnse("expired token", err.Error(), helper.EmptyObj{})
				return c.JSONPretty(http.StatusUnauthorized, res, "  ")
			}
		}

		c.Set("email", claims["email"].(string))

		return next(c)
	}
}
