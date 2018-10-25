package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Participants struct dto
type Participants struct {
	UuidParticipants     uuid.UUID
	ParticipantsName     string
	ParticipantsAddress  string
	ParticipantsBirthday time.Time
	ParticipantsAge      int
	ParticipantsGender   string
	ParticipantsPhone    int
}
