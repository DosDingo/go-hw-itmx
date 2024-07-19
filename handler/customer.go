package handler

import (
	"goassignment/database"
	"goassignment/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCustomer(c echo.Context) error {
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	database.DB.Create(customer)
	return c.JSON(http.StatusCreated, customer)
}

func GetCustomers(c echo.Context) error {
	var customers []model.Customer
	database.DB.Find(&customers)
	return c.JSON(http.StatusOK, customers)
}

func GetCustomer(c echo.Context) error {
	id := c.Param("id")
	var customer model.Customer
	result := database.DB.First(&customer, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Customer not found"})
	}
	return c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c echo.Context) error {
	id := c.Param("id")
	var customer model.Customer
	result := database.DB.First(&customer, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Customer not found"})
	}

	if err := c.Bind(&customer); err != nil {
		return err
	}

	database.DB.Save(&customer)
	return c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	var customer model.Customer
	result := database.DB.Delete(&customer, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Customer not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
