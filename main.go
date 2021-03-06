package main

import (
	"codemi/controller"

	"github.com/gin-gonic/gin"
)

// MAIN APPLICATION
func main() {
	BaseRoutingApplication()
}

// BaseRoutingApplication func does store all routing application
func BaseRoutingApplication() {
	baseRoutes := gin.Default()

	// Versioning API
	v1 := baseRoutes.Group("/v1/api/")
	{
		// Participants base routing
		v1.GET("participants", controller.GetAllParticipants)
		v1.GET("participants/:UuidParticipants", controller.GetDetailParticipants)
		v1.POST("participants", controller.CreateNewParticipants)

		// Classroom base routing
		v1.GET("classroom", controller.GetAllClass)
		v1.GET("classroom/:UuidClassroom", controller.GetDetailClass)
		v1.POST("classroom", controller.CreateNewClass)
	}

	baseRoutes.Run(":8000")
}
