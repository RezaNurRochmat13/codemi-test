package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Participants struct model
type Participants struct {
	gorm.Model
	IdParticipants       uuid.UUID
	UuidParticipants     uuid.UUID
	ParticipantsName     string
	ParticipantsAddress  string
	ParticipantsBirthday time.Time
	ParticipantsAge      int
	ParticipantsGender   string
	ParticipantsPhone    int
}
