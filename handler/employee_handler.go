package handler

import (
	"belajar_rest_api_pribadi/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllEmployee(c echo.Context) (err error) {
	c.JSON(http.StatusOK, entities.Employee{
		Id: 1,
		Name: "Krisna Satriadi",
		Position: "TEC",
	})
	return
}