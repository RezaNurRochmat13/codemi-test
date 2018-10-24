package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Participants struct
type Participants struct {
	gorm.Model
	IDParticipants      uuid.UUID
	UUIDParticipants    uuid.UUID
	ParticipantName     string
	ParticipantAddress  string
	ParticipantBirthday string
	ParticipantAge      int
	ParticipantGender   string
	ParticipantPhone    string
}
