package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


// GetAllClass func does get all class
func GetAllClass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": 1,
		"data": "All data class"
		"count": 1
	})
}

// CreateNewClass func does create new class and participants
func CreateNewClass(c *gin.Context) {
	c.JSON("message" : "Inserted successfully")
}