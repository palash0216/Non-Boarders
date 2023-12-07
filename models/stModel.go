package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string
	Enroll  string
	Place   string
	TimeIn  time.Time
	TimeOut time.Time
}
