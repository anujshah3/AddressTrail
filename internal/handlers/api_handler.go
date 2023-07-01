package handlers

import (
	"net/http"
	"time"

	"github.com/anujshah3/AddressTrail/internal/models"
	"github.com/anujshah3/AddressTrail/internal/services"
	"github.com/gin-gonic/gin"
)

func AddNewUserHandler(c *gin.Context) {
	user := &models.User{
		ID:    "",
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
		Addresses: []*models.AddressWithDates{},
	}
	userID, err := services.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User inserted successfully",
		"user":    userID,
	})
}

func DeleteUserHandler(c *gin.Context) {
	userID := c.PostForm("userID")
	err := services.DeleteUser(userID)
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

type AddAddressToUserPayload struct {
	UserID     string
	Street     string
	Unit       string
	City       string
	State      string
	PostalCode string
	Country    string
	StartDate  string
	EndDate    string
}

func AddAddressToUserHandler(c *gin.Context) {
    var payload AddAddressToUserPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Payload",
		})
		return
	}

	userID := payload.UserID
	startDateStr := payload.StartDate
	endDateStr := payload.EndDate
	
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

	address := &models.Address{
		Street:     payload.Street,
		Unit:       payload.Unit,
		City:       payload.City,
		State:      payload.State,
		PostalCode: payload.PostalCode,
		Country:    payload.Country,
	}

	addressID, err := services.InsertAddress(address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	addressWithDates := &models.AddressWithDates{
		AddressID: addressID,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = services.AddNewAddressToUser(userID, addressWithDates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"addressID": addressID,
		"message": "Address added to user successfully",
	})
}


func UpdateUserAddressHandler(c *gin.Context) {
	userID := c.PostForm("userID")
	addressID := c.PostForm("addressID")

	address, err := services.GetAddressByID(addressID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Address not found",
		})
		return
	}

	// Check if fields are present
	fieldsUpdated := false

	if street := c.PostForm("street"); street != "" {
		address.Street = street
		fieldsUpdated = true
	}
	if unit := c.PostForm("unit"); unit != "" {
		address.Unit = unit
		fieldsUpdated = true
	}
	if city := c.PostForm("city"); city != "" {
		address.City = city
		fieldsUpdated = true
	}
	if state := c.PostForm("state"); state != "" {
		address.State = state
		fieldsUpdated = true
	}
	if postalCode := c.PostForm("postalCode"); postalCode != "" {
		address.PostalCode = postalCode
		fieldsUpdated = true
	}
	if country := c.PostForm("country"); country != "" {
		address.Country = country
		fieldsUpdated = true
	}

	if fieldsUpdated {
		err = services.UpdateAddress(addressID, address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	oldStartDateStr := c.PostForm("oldStartDate")
	oldEndDateStr := c.PostForm("oldEndDate")
	newStartDateStr := c.PostForm("newStartDate")
	newEndDateStr := c.PostForm("newEndDate")

	if oldStartDateStr != "" && oldEndDateStr != "" && newStartDateStr != "" && newEndDateStr != "" {
		oldStartDate, err := time.Parse("2006-01-02", oldStartDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid oldStartDate format",
			})
			return
		}

		oldEndDate, err := time.Parse("2006-01-02", oldEndDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid oldEndDate format",
			})
			return
		}

		newStartDate, err := time.Parse("2006-01-02", newStartDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid newStartDate format",
			})
			return
		}

		newEndDate, err := time.Parse("2006-01-02", newEndDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid newEndDate format",
			})
			return
		}

		err = services.UpdateFilteredAddresses(userID, addressID, oldStartDate, oldEndDate, newStartDate, newEndDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Address updated successfully",
	})
}

// func DeleteAddressFromUserHandler(c *gin.Context) {
// 	userID := c.PostForm("userID")
// 	addressID := c.PostForm("addressID")
// 	startDateStr := c.PostForm("startDate")
// 	endDateStr := c.PostForm("endDate")

// 	startDate, err := time.Parse("2006-01-02", startDateStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid startDate format",
// 		})
// 		return
// 	}

// 	endDate, err := time.Parse("2006-01-02", endDateStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid endDate format",
// 		})
// 		return
// 	}

// 	address := &models.AddressWithDates{
// 		AddressID: addressID,
// 		StartDate: startDate,
// 		EndDate:   endDate,
// 	}

// 	err = services.DeleteAddressFromUser(userID, address)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Address deleted from user successfully",
// 	})
// }


func GetUserAddressesHandler(c *gin.Context) {
	userID := c.Query("userID")
	addresses, err := services.GetUserAddresses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, addresses)
}
