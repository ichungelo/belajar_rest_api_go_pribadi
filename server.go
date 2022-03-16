package main

import (
	"belajar_rest_api_pribadi/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
    e.Use(middleware.CORS())

    employee := e.Group("/api/employee")
    {
        employee.GET("", handler.GetAllEmployee)
    }
	// shift := e.Group("/api/shift")
    e.Logger.Fatal(e.Start(":8080"))
}
