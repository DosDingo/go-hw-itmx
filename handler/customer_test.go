package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"goassignment/database"
	"goassignment/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error
	database.DB, err = gorm.Open(sqlite.Open("test_customers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.DB.AutoMigrate(&model.Customer{})
}

func TestCreateCustomer(t *testing.T) {
	setupTestDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"Alice", "age": 29}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateCustomer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), `"name":"Alice"`)
	}
}

func TestGetCustomers(t *testing.T) {
	setupTestDB()
	e := echo.New()

	// Create customers first
	database.DB.Create(&model.Customer{Name: "Alice", Age: 29})
	database.DB.Create(&model.Customer{Name: "Bob", Age: 34})

	req := httptest.NewRequest(http.MethodGet, "/customers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetCustomers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"name":"Alice"`)
		assert.Contains(t, rec.Body.String(), `"name":"Bob"`)
	}
}

func TestGetCustomer(t *testing.T) {
	setupTestDB()
	e := echo.New()

	// Create a customer first
	customer := model.Customer{Name: "Bob", Age: 34}
	database.DB.Create(&customer)

	req := httptest.NewRequest(http.MethodGet, "/customers/"+fmt.Sprint(customer.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customers/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(customer.ID))

	if assert.NoError(t, GetCustomer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"name":"Bob"`)
	}
}

func TestUpdateCustomer(t *testing.T) {
	setupTestDB()
	e := echo.New()

	// Create a customer first
	customer := model.Customer{Name: "Charlie", Age: 28}
	database.DB.Create(&customer)

	req := httptest.NewRequest(http.MethodPut, "/customers/"+fmt.Sprint(customer.ID), strings.NewReader(`{"name":"Charlie Updated", "age": 29}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customers/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(customer.ID))

	if assert.NoError(t, UpdateCustomer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"name":"Charlie Updated"`)
	}
}

func TestDeleteCustomer(t *testing.T) {
	setupTestDB()
	e := echo.New()

	// Create a customer first
	customer := model.Customer{Name: "David", Age: 40}
	database.DB.Create(&customer)

	req := httptest.NewRequest(http.MethodDelete, "/customers/"+fmt.Sprint(customer.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customers/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(customer.ID))

	if assert.NoError(t, DeleteCustomer(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
