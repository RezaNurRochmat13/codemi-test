package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Participants struct mapper
type Participants struct {
	UuidParticipants     uuid.UUID
	ParticipantsName     string
	ParticipantsAddress  string
	ParticipantsBirthday time.Time
	ParticipantsAge      string
	ParticipantsGender   string
	ParticipantsPhone    string
}
