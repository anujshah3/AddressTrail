package main

import (
	"github.com/anujshah3/AddressTrail/internal/routes"
	"github.com/gin-gonic/gin"
)



func main() {
    router := gin.Default()

	routes.SetupDevAPIRoutes(router)
	routes.SetupAPIRoutes(router)

	routes.SetupWebRoutes(router)

	router.Run(":8080")

}