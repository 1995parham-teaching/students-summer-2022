package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello")
	})

	if err := app.Start(":1234"); err != nil {
		log.Println(err)
	}
}
