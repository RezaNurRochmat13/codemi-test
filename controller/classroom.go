package controller

import (
	"codemi/config"
	"codemi/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* @created by Reja Nur Rochmat
* @email rezanurrochmat3@gmail.com
**/

// GetAllClass func does get all class
func GetAllClass(c *gin.Context) {

	// Initialize database connection
	db := config.DatabaseConn()

	// Declaring models dto
	var (
		rowsCount     int
		dataClassroom []dto.Classroom
	)

	// Get all record from database
	db.Find(&dataClassroom)

	// Counting record from database
	db.Find(&dataClassroom).Count(&rowsCount)

	// Serve as JSON
	c.JSON(http.StatusOK, gin.H{
		"total": rowsCount,
		"data":  dataClassroom,
		"count": rowsCount})
}

// GetDetailClass func does get detail classroom
func GetDetailClass(c *gin.Context) {
	// Initialize database connection
	db := config.DatabaseConn()

	// Declare UUID parameters
	UuidClassroom := c.Param("UuidClassroom")

	// Declare models dto
	var detailClassroom dto.Classroom

	// Find the record in database
	db.Where("classrooms.uuid_classroom = ?", UuidClassroom).Find(&detailClassroom)

	// Checking in database
	if db.Where("classrooms.uuid_classroom = ?", UuidClassroom).Find(&detailClassroom).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": detailClassroom})
	}
}

// CreateNewClass func does create new class and participants
func CreateNewClass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}
