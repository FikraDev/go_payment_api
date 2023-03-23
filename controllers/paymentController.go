package controllers

import (
	"errors"
	"net/http"
	"paymentapi/initializers"
	"paymentapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// paymentFound := false

	// paymentNotFound := true

	// var payment []models.Payment
	// initializers.DB.First(&payment, id)

	// if paymentFound {
	// 	c.JSON(http.StatusOK, payment)
	// 	paymentNotFound = false
	// }

	// if paymentNotFound {

	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"code":  "404",
	// 		"error": "Payment Not Found",
	// 	})
	// 	return
	// }

	// c.JSON(200, gin.H{
	// 	"payment": payment,
	// })
	var payment []models.Payment
	if err := initializers.DB.First(&payment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Payment Not Found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  "500",
				"error": "Internal Server Error",
			})
		}
		return
	}

	// payment was found, so return it
	c.JSON(http.StatusOK, payment)

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
