package controller

import (
	"codemi/config"
	"codemi/model"
	"codemi/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
	var detailClassroom dto.Classrooms

	// Custom select columns
	db.Select("classrooms.uuid_classroom, classrooms.classroom_name, classrooms.classroom_time, classrooms.room, participants.participants_name, participants.participants_address, participants.participants_gender, participants.participants_phone").
		Joins("INNER JOIN participants ON participants.uuid_participants=classrooms.uuid_participants").
		Where("classrooms.uuid_classroom = ?", UuidClassroom).
		Find(&detailClassroom)

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

	// // Initialize database connection
	db := config.DatabaseConn()

	// Declare parameters
	IDClassroom := uuid.Must(uuid.NewV4())
	UUIDClassroom := uuid.Must(uuid.NewV4())
	ClassroomName := c.Param("ClassroomName")
	ClassroomTime := c.Param("ClassroomTime")
	Room := c.Param("Room")
	UUIDParticipants := c.Param("UUIDParticipants")

	// Bind in one parameters
	createClassroomPayload := model.Classroom{
		IDClassroom:      IDClassroom,
		UUIDClassroom:    UUIDClassroom,
		ClassroomName:    ClassroomName,
		ClassroomTime:    ClassroomTime,
		Room:             Room,
		UUIDParticipants: UUIDParticipants}

	// Bind parameter as JSON Parameters
	c.BindJSON(&createClassroomPayload)

	// Save in database
	db.Save(&createClassroomPayload)

	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}
