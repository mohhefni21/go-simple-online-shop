package middleware

import (
	"mohhefni/go-online-shop/utility"

	"github.com/labstack/echo/v4"
)

func Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utility.MakeLogEntry(c).Info("Incoming Request")
		return next(c)
	}
}
