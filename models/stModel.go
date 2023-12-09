package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	Enroll string
	Place  string
	In     string
	Out    string
}
