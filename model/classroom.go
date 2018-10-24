package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Classroom struct
type Classroom struct {
	gorm.Model
	IDClassroom   uuid.UUID
	UUIDClassroom uuid.UUID
	ClassroomName string
	ClassroomTime string
	Room          string
	Attendees     []Participants
}
