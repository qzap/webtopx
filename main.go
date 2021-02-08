package main

import (
	"net/http"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.Get("/", func(c echo.Context) error {
		return c.String("Hello, World!", http.StatusOK)
	})
	e.Run(standard.New(":1323"))
}
