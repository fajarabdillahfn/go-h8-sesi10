package controllers

import (
	"fmt"
	"net/http"
	"sesi10/database"
	"sesi10/helpers"
	"sesi10/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	var product models.Product

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	userID := uint(userData["id"].(float64))
	product.UserID = userID

	err := db.Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	fmt.Printf("c.Request.Body: %v\n", c.Request.Body)
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	var product models.Product

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	fmt.Printf("product: %v\n", product)
	productID, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid parameter",
		})
		return
	}
	userID := uint(userData["id"].(float64))

	product.UserID = userID
	product.ID = uint(productID)

	err = db.Model(&models.Product{}).Where("id = ?", productID).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
