package dto

import (
	uuid "github.com/satori/go.uuid"
)

// Classroom struct
type Classroom struct {
	UUIDClassroom uuid.UUID
	ClassroomName string
	ClassroomTime string
	Room          string
	Attendees     Participants `gorm:"many2many:uuid_participants;"`
}
