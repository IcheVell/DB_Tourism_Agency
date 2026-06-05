package middleware

import (
	"net/http"
	"strings"

	"TouristAgencyApp/src/auth"

	"github.com/labstack/echo/v4"
)

const ContextUserID = "user_id"
const ContextUserLogin = "user_login"

func JWTAuth(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "authorization header is required",
				})
			}

			const bearerPrefix = "Bearer "

			if !strings.HasPrefix(authHeader, bearerPrefix) {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "authorization header must be Bearer token",
				})
			}

			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, bearerPrefix))
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "token is required",
				})
			}

			claims, err := auth.ParseAccessToken(tokenString, secret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired token",
				})
			}

			c.Set(ContextUserID, claims.UserID)
			c.Set(ContextUserLogin, claims.Login)

			return next(c)
		}
	}
}

func GetUserID(c echo.Context) (int64, bool) {
	value := c.Get(ContextUserID)
	if value == nil {
		return 0, false
	}

	userID, ok := value.(int64)
	return userID, ok
}
