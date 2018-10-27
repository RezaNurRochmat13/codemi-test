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

// GetAllParticipants func does get all participants
func GetAllParticipants(c *gin.Context) {
	// Initialize database connection
	db := config.DatabaseConn()

	// Declare variable
	var (
		countParticipants int
		dataParticipants  []dto.Participants
	)

	// Get all data from database
	db.Find(&dataParticipants)

	// Counting row in database
	db.Find(&dataParticipants).Count(&countParticipants)

	// Serve as JSON
	c.JSON(http.StatusOK, gin.H{
		"count": countParticipants,
		"data":  dataParticipants,
		"total": countParticipants})
}

// GetDetailParticipants func does get detail participants
func GetDetailParticipants(c *gin.Context) {
	// Initialize database connection
	db := config.DatabaseConn()

	// Declare variable and model dto
	var detailParticipants dto.Participants

	// Parameter UUID Participants
	UuidParticipants := c.Param("UuidParticipants")

	// Find record in database by uuid participants
	db.Where("participants.uuid_participants = ?", UuidParticipants).Find(&detailParticipants)

	// Checking record in database
	if db.Where("participants.uuid_participants = ?", UuidParticipants).Find(&detailParticipants).RecordNotFound() {
		// When record not found, returned 404 Not Found
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
	} else {
		// When record found, returned 200 OK and data participants
		c.JSON(http.StatusOK, gin.H{"data": detailParticipants})
	}
}

// CreateNewParticipants func does create new participants
func CreateNewParticipants(c *gin.Context) {
	// Initialize database connection
	db := config.DatabaseConn()

	// Declaring input parameters
	UuidParticipants := uuid.Must(uuid.NewV4())
	ParticipantsName := c.Param("ParticipantsName")
	ParticipantsAddress := c.Param("ParticipantsAddress")
	ParticipantsBirthday := c.Param("ParticipantsBirthday")
	ParticipantsAge := c.Param("ParticipantsAge")
	ParticipantsGender := c.Param("ParticipantsGender")
	ParticipantsPhone := c.Param("ParticipantsPhone")

	// Bind in one parameters
	createParticipantsPayload := model.Participants{
		UuidParticipants:     UuidParticipants,
		ParticipantsName:     ParticipantsName,
		ParticipantsAddress:  ParticipantsAddress,
		ParticipantsBirthday: ParticipantsBirthday,
		ParticipantsAge:      ParticipantsAge,
		ParticipantsGender:   ParticipantsGender,
		ParticipantsPhone:    ParticipantsPhone}

	// Bind as JSON parameters
	c.BindJSON(&createParticipantsPayload)

	// Saved in database
	db.Save(&createParticipantsPayload)

	// Success message
	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}
