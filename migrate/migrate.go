package main

import (
	"paymentapi/initializers"
	"paymentapi/models"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnection()
}

func main() {

	//migrate the schema
	initializers.DB.AutoMigrate(&models.Payment{})

}
