package routes

import (
	"github.com/anujshah3/AddressTrail/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	router.DELETE("/api/users", handlers.DeleteUserHandler)
	router.POST("/api/users/addresses", handlers.AddAddressToUserHandler)
	router.PATCH("/api/users/addresses", handlers.DeleteAddressFromUserHandler)
	router.GET("/api/users/addresses", handlers.GetUserAddressesHandler)
}
