package main

import (
	"paymentapi/controllers"
	"paymentapi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnection()
}

func main() {

	r := gin.Default()

	//routes

	r.POST("/payment", controllers.AddPayment)       //create
	r.PUT("/payment/:id", controllers.UpdatePayment) //create

	r.GET("/allpayments", controllers.GetAllPayments)   //get all payments
	r.GET("/payment/:id", controllers.GetSinglePayment) //get single payment
	r.DELETE("/payment/:id", controllers.DeletePayment) //get single payment

	r.Run(":3000") // listen and serve on 0.0.0.0:3000 calling the router

}
