package handlers

import (
	"net/http"
	"time"

	"github.com/anujshah3/AddressTrail/internal/models"
	"github.com/anujshah3/AddressTrail/internal/services"
	"github.com/gin-gonic/gin"
)


func DeleteUserHandler(c *gin.Context) {
	email := c.PostForm("email")
	err := services.DeleteUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}


func AddAddressToUserHandler(c *gin.Context) {
	userID := c.PostForm("userID")
	addressID := c.PostForm("addressID")
	startDateStr := c.PostForm("startDate")
	endDateStr := c.PostForm("endDate")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid startDate format",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid endDate format",
		})
		return
	}

	address := &models.AddressWithDates{
		AddressID: addressID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	err = services.AddNewAddressToUser(userID, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Address added to user successfully",
	})
}


func DeleteAddressFromUserHandler(c *gin.Context) {
	userID := c.PostForm("userID")
	addressID := c.PostForm("addressID")
	startDateStr := c.PostForm("startDate")
	endDateStr := c.PostForm("endDate")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid startDate format",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid endDate format",
		})
		return
	}

	address := &models.AddressWithDates{
		AddressID: addressID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	err = services.DeleteAddressFromUser(userID, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Address deleted from user successfully",
	})
}


func GetUserAddressesHandler(c *gin.Context) {
	email := c.PostForm("email")
	addresses, err := services.GetUserAddresses(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}
