package middleware

import (
	"mohhefni/go-online-shop/infra/errorpkg"
	"mohhefni/go-online-shop/infra/responsepkg"
	"mohhefni/go-online-shop/utility"

	"github.com/labstack/echo/v4"
)

func CheckRole(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			roleUser := c.Get("role").(string)

			isExits := false
			for _, role := range roles {
				if role == roleUser {
					isExits = true
					break
				}
			}

			if !isExits {
				err := errorpkg.ErrForbiddenAccess
				utility.MakeLogEntry(nil).Warning(err)
				return responsepkg.NewResponse(
					responsepkg.WithStatus(err),
				).Send(c)
			}

			return next(c)

		}
	}
}
