package main

import (
	"myapp/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func publicHandler(c echo.Context) error {
	// User ID from path `users/:id`
	return c.String(http.StatusOK, "Hello, World!")
}

func privateHandler(c echo.Context) error {
	// User ID from path `users/:id`
	return c.String(http.StatusOK, "Hello, Admin!")
}

func main() {

	e := echo.New()
	e.GET("/", publicHandler)

	//Restricted paths
	r := e.Group("/admin", middlewares.BasicAuth())

	r.GET("", privateHandler)
	e.Logger.Fatal(e.Start(":8000"))

}
