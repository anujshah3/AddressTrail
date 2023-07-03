package main

import (
	"os"

	"github.com/anujshah3/AddressTrail/internal/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)



func main() {
    router := gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	router.Use(sessions.Sessions("session", store))

	routes.SetupDevAPIRoutes(router)
	routes.SetupAPIRoutes(router)

	routes.SetupWebRoutes(router)

	router.Run(":8080")

}