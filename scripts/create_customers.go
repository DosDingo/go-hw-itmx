package main

import (
	"goassignment/database"
	"goassignment/model"
	"log"
)

func main() {
	database.InitDatabase()

	database.DB.Where("1 = 1").Delete(&model.Customer{})

	customers := []model.Customer{
		{Name: "John Doe", Age: 20},
		{Name: "Jane Doe", Age: 25},
		{Name: "Jame Brown", Age: 30},
		{Name: "John Bame", Age: 35},
		{Name: "Jame Born", Age: 35},
	}

	for _, customer := range customers {
		database.DB.Create(&customer)
	}

	log.Println("Customers created successfully")
}
