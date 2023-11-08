package controllers

import (
	"phoneService/service"

	"github.com/labstack/echo/v4"
)

func GETAllPhone(c echo.Context) error {

	term := c.QueryParam("term")

	phones, err := service.New().GetAll(term)

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, phones)

}
