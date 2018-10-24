package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* @created by Reja Nur Rochmat
* @email rezanurrochmat3@gmail.com
**/

// GetAllParticipants func does get all participants
func GetAllParticipants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count": 1,
		"data":  "All data participants",
		"total": 1})
}

// GetDetailParticipants func does get detail participants
func GetDetailParticipants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Detail data participants"})
}

// CreateParticipants func does create new participants
func CreateParticipants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}
