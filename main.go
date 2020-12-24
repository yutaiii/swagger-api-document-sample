package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @contact.name Example API Support
// @contact.url http://example.com/support
// @contact.email support@example.com
// @BasePath /
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routing
	e.GET("/hello-world", helloWorld())

	//exec server
	e.Start(":8000")
}

// @Summary hello world
// @Description say hello world
// @ID hello-world
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /hello-world [get]
func helloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	}
}
