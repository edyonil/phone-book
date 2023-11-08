package main

import (
	"net/http"
	"phoneService/config"
	"phoneService/controllers"

	"github.com/labstack/echo/v4"
)

func main() {

	err := config.InitDB()

	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/search", controllers.GETAllPhone)
	e.Logger.Fatal(e.Start(":1323"))
}
