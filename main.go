package main

import (
	"goassignment/database"
	"goassignment/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDatabase()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/customers", handler.CreateCustomer)
	e.GET("/customers", handler.GetCustomers)
	e.GET("/customers/:id", handler.GetCustomer)
	e.PUT("/customers/:id", handler.UpdateCustomer)
	e.DELETE("/customers/:id", handler.DeleteCustomer)

	e.Logger.Fatal(e.Start(":8080"))
}
