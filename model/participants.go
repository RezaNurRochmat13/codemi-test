package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Participants struct model
type Participants struct {
	gorm.Model
	UuidParticipants     uuid.UUID
	ParticipantsName     string
	ParticipantsAddress  string
	ParticipantsBirthday string
	ParticipantsAge      string
	ParticipantsGender   string
	ParticipantsPhone    string
}
