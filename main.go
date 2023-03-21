package main

import (
	"net/http"
	"paymentapi/controllers"
	"paymentapi/initializers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnection()
}

type Payment struct {
	gorm.Model
	Amount      float64
	Quantity    float64
	Description string
	ItemCode    string
}

func main() {

	r := gin.Default()

	//routes

	//r.POST("/payment", controllers.AddPayment)       //create

	r.POST("/payment", func(c *gin.Context) {

		amount := c.PostForm("amount")
		quantity := c.PostForm("quantity")
		description := c.PostForm("description")
		itemcode := c.PostForm("itemcode")

		payment := Payment{Amount: c.GetFloat64(amount), Quantity: c.GetFloat64(quantity), Description: description, ItemCode: itemcode}

		initializers.DB.Create(&payment)

		c.JSON(http.StatusOK, gin.H{
			"message": "Data submitted successfully",
		})

	})

	r.PUT("/payment/:id", controllers.UpdatePayment) //create

	r.GET("/allpayments", controllers.GetAllPayments)   //get all payments
	r.GET("/payment/:id", controllers.GetSinglePayment) //get single payment
	r.DELETE("/payment/:id", controllers.DeletePayment) //get single payment

	r.Run(":3000") // listen and serve on 0.0.0.0:3000 calling the router

}
