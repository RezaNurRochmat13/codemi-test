package main

import (
	"codemi/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	BaseRoutingApplication()
}

// BaseRoutingApplication func does store all routing application
func BaseRoutingApplication() {
	baseRoutes := gin.Default()

	// Versioning API
	v1 := baseRoutes.Group("/v1/api/")
	{
		v1.GET("participants", controller.GetAllParticipants)
		v1.GET("participants/:UuidParticipants", controller.GetDetailParticipants)
	}

	baseRoutes.Run(":8000")
}
