package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID
	Login          string
	HashedPassword string
	LastSeen       time.Time
}

type Task struct {
	Id              uuid.UUID
	UserId          uuid.UUID
	TaskName        string
	TaskDescription string
	ImageUrls       string
}
