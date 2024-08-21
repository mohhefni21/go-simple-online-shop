package utility

import (
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		ForceColors:   true,
	})
}

func MakeLogEntry(ctx echo.Context) *log.Entry {
	if ctx == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": ctx.Request().Method,
		"uri":    ctx.Request().URL.String(),
		"ip":     ctx.Request().RemoteAddr,
	})
}
