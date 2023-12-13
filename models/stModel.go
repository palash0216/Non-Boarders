package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `jsno:"name"`
	Enroll string `jsno:"enroll"`
	Place  string `jsno:"place"`
	In     string `jsno:"in"`
	Out    string `jsno:"out"`
}
