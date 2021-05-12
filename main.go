package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello Echo Framework!!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
