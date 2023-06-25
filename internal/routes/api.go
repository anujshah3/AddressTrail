package routes

import (
	"github.com/anujshah3/AddressTrail/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	router.POST("/api/users", handlers.AddNewUserHandler)
	router.DELETE("/api/users", handlers.DeleteUserHandler)

	router.GET("/api/users/addresses", handlers.GetUserAddressesHandler)
	router.POST("/api/users/addresses", handlers.AddAddressToUserHandler)
	router.PUT("/api/users/addresses", handlers.UpdateUserAddressHandler)
	// router.PATCH("/api/users/addresses", handlers.DeleteAddressFromUserHandler)
}
