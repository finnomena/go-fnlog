// demo with echo
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/finnomena/go-fnlog"
)

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(fnlog.LoggingMiddleware()))
	e.GET("/", func(c echo.Context) error {
		fnlog.LogWithoutContext(fnlog.ErrorLevel).Info("test")
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
