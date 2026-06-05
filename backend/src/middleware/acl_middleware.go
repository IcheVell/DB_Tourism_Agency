package middleware

import (
	"net/http"

	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

func RequirePermission(aclService *service.ACLService, permissionCode string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, ok := GetUserID(c)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "user is not authenticated",
				})
			}

			hasPermission, err := aclService.HasPermission(userID, permissionCode)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "failed to check permission",
				})
			}

			if !hasPermission {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "permission denied",
				})
			}

			return next(c)
		}
	}
}
