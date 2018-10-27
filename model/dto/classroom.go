package dto

import (
	uuid "github.com/satori/go.uuid"
)

// Classroom struct list mapper
type Classroom struct {
	UUIDClassroom uuid.UUID
	ClassroomName string
	ClassroomTime string
	Room          string
}

// Classrooms struct detail mapper
type Classrooms struct {
	UUIDClassroom       uuid.UUID
	ClassroomName       string
	ClassroomTime       string
	Room                string
	ParticipantsName    string
	ParticipantsAddress string
	ParticipantsGender  string
	ParticipantsPhone   string
}
