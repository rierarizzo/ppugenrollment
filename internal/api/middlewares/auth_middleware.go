package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"ppugenrollment/pkg/domain"
	"ppugenrollment/pkg/security"
	"strings"

	"github.com/labstack/echo/v4"
)

var RoutesAllowedByRoles map[string][]string

func VerifyJWTAndRoles(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validate JWT
		tokenWithBearer := c.Request().Header.Get("Authorization")

		if tokenWithBearer == "" {
			slog.Error("Token is empty")
			return c.JSON(http.StatusForbidden, domain.NewAppError("Token is empty", domain.TokenValidationError))
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := security.VerifyJWTToken(token)

		if err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusForbidden, domain.NewAppError(err.Error(), domain.TokenValidationError))
		}

		c.Set("UserID", claims.Id)

		// Validate roles
		var rolesAllowed []string

		for k, v := range RoutesAllowedByRoles {
			if c.Path() == k {
				rolesAllowed = v
				break
			}

			if strings.HasPrefix(c.Path(), k) {
				rolesAllowed = v
			}
		}

		if rolesAllowed[0] == "ALL" {
			return next(c)
		}

		for _, v := range rolesAllowed {
			if claims.Role == v {
				return next(c)
			}
		}

		slog.Error(fmt.Sprintf("Role %s not authorized", claims.Role))

		return c.JSON(http.StatusUnauthorized, domain.NewAppErrorWithType(domain.NotAuthorizedError))
	}
}
