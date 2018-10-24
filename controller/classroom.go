package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* @created by Reja Nur Rochmat
* @email rezanurrochmat3@gmail.com
**/

// GetAllClass func does get all class
func GetAllClass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": 1,
		"data":  "All data class",
		"count": 1})
}

// CreateNewClass func does create new class and participants
func CreateNewClass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}

// DetailClass func does get detail classroom
func DetailClass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Detail classroom"})
}
