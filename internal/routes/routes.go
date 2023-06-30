package routes

import (
	"net/http"

	"github.com/anujshah3/AddressTrail/internal/handlers"
	"github.com/gin-gonic/gin"
)
	func SetupDevAPIRoutes(router *gin.Engine) {
		router.POST("/api/dev/users", handlers.AddNewUserHandler)
		router.DELETE("/api/dev/users", handlers.DeleteUserHandler)

		router.GET("/api/dev/users/addresses", handlers.GetUserAddressesHandler)
		router.POST("/api/dev/users/addresses", handlers.AddAddressToUserHandler)
		router.PUT("/api/dev/users/addresses", handlers.UpdateUserAddressHandler)
		// router.PATCH("/api/dev/users/addresses", handlers.DeleteAddressFromUserHandler)
	}

	func SetupAPIRoutes(router *gin.Engine) {
		router.POST("/api/users", handlers.AddNewUserHandler)
		router.DELETE("/api/users", handlers.DeleteUserHandler)

		router.GET("/api/users/addresses", handlers.GetUserAddressesHandler)
		router.POST("/api/users/addresses", handlers.AddAddressToUserHandler)
		router.PUT("/api/users/addresses", handlers.UpdateUserAddressHandler)
		// router.PATCH("/api/users/addresses", handlers.DeleteAddressFromUserHandler)
	}

	func SetupWebRoutes(router *gin.Engine) {
		router.LoadHTMLGlob("web/templates/*.html")
		router.GET("/", func(c *gin.Context) {
        	c.HTML(http.StatusOK, "index.html", nil)
    	})
		router.GET("/login", gin.WrapF(handlers.GoogleLoginHandler))
		// router.GET("/auth/google/callback", gin.WrapF(handlers.GoogleCallBackHandler))
		// router.GET("/dashboard", gin.WrapF(middleware.AuthMiddleware(handlers.DashboardHandler)))
		// router.GET("/address-book", gin.WrapF(middleware.AuthMiddleware(handlers.AddressBookHandler)))

		// router.GET("/logout", gin.WrapF(handlers.LogoutHandler))

		router.Static("/static", "./web/static")
	}
