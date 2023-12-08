package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	initializers "github.com/palash0216/Non-Boarders/Initializers"
	"github.com/palash0216/Non-Boarders/models"
)

func StCreate(c *gin.Context) {
	//Getting data
	var body struct {
		Name    string
		Enroll  string
		Place   string
		TimeIn  time.Time
		TimeOut time.Time
	}
	c.Bind(&body)

	//Creating a Student
	student := models.Student{Name: body.Name, Enroll: body.Enroll, Place: body.Place, TimeIn: body.TimeIn, TimeOut: body.TimeOut}
	result := initializers.DB.Create(&student)

	if result.Error != nil {
		c.Status(404)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"student": student,
	})
}
