package controllers

import (
	"paymentapi/initializers"
	"paymentapi/models"

	"github.com/gin-gonic/gin"
)

// func AddPayment(c *gin.Context) {

// 	var Payment struct {
// 		Amount      float64
// 		Quantity    float64
// 		Description string
// 		ItemCode    string
// 	}

// 	c.Bind(&Payment)

// 	payment := models.Payment{Amount: Payment.Amount, Quantity: Payment.Quantity, Description: Payment.Description, ItemCode: Payment.ItemCode}
// 	if err := c.ShouldBindJSON(&payment); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result := initializers.DB.Create(&payment)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": payment})
// }

//get all payments

func GetAllPayments(c *gin.Context) {

	var payments []models.Payment
	initializers.DB.Find(&payments)

	c.JSON(200, gin.H{
		"payment": payments,
	})

}

func GetSinglePayment(c *gin.Context) {

	id := c.Param("id")

	var payment []models.Payment
	initializers.DB.First(&payment, id)

	c.JSON(200, gin.H{
		"payment": payment,
	})

}

//Update

func UpdatePayment(c *gin.Context) {

	id := c.Param("id")

	var Payment struct {
		Amount      float64
		Quantity    float64
		Description string
		ItemCode    string
	}

	c.Bind(&Payment)

	var payment []models.Payment
	initializers.DB.First(&payment, id)

	initializers.DB.Model(&payment).Updates(models.Payment{

		Amount:      Payment.Amount,
		Quantity:    Payment.Quantity,
		Description: Payment.Description,
		ItemCode:    Payment.ItemCode,
	})

	c.JSON(200, gin.H{
		"payment": payment,
	})

}

//Delete

func DeletePayment(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Payment{}, id)

	c.Status(200)

}
