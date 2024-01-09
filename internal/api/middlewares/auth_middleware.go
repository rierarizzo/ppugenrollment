package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"ppugenrollment/pkg/domain"
	"ppugenrollment/pkg/security"
	"ppugenrollment/pkg/utils"
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
			return utils.SendError(http.StatusForbidden,
				domain.NewAppError("Token is empty", domain.TokenValidationError))
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := security.VerifyJWTToken(token)

		if err != nil {
			slog.Error(err.Error())
			return utils.SendError(http.StatusForbidden, domain.NewAppErrorWithType(domain.TokenValidationError))
		}

		c.Set("UserID", claims.Id)

		// Validate roles
		isValid := validateRoles(c.Path(), claims.Role)

		if isValid {
			return next(c)
		}

		appErrMsg := fmt.Sprintf("Role %s not authorized", claims.Role)
		slog.Error(appErrMsg)

		return utils.SendError(http.StatusUnauthorized, domain.NewAppError(appErrMsg, domain.NotAuthorizedError))
	}
}

func validateRoles(path string, role string) bool {
	var rolesAllowed []string

	for k, v := range RoutesAllowedByRoles {
		if path == k {
			rolesAllowed = v
			break
		}

		if strings.HasPrefix(path, k) {
			rolesAllowed = v
		}
	}

	if rolesAllowed[0] == "ALL" {
		return true
	}

	for _, v := range rolesAllowed {
		if role == v {
			return true
		}
	}

	return false
}
