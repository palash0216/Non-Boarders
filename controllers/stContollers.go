package controllers

import (
	"github.com/gin-gonic/gin"
	initializers "github.com/palash0216/Non-Boarders/Initializers"
	"github.com/palash0216/Non-Boarders/models"
)

func StCreate(c *gin.Context) {
	//Getting data
	var body struct {
		Name   string
		Enroll string
		Place  string
		In     string
		Out    string
	}
	c.Bind(&body)

	//Creating a Student
	student := models.Student{Name: body.Name, Enroll: body.Enroll, Place: body.Place, In: body.In, Out: body.Out}
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

func StudentIndex(c *gin.Context) {
	//Get the post
	var students []models.Student
	initializers.DB.Find(&students)
	//repond to them
	c.JSON(200, gin.H{
		"students": students,
	})
}

func StudentShow(c *gin.Context) {
	//Get ID of URL
	id := c.Param("id")
	//Get the post
	var student models.Student
	initializers.DB.First(&student, id)
	//repond to them
	c.JSON(200, gin.H{
		"student": student,
	})
}

func StUpdate(c *gin.Context) {
	//Get ID off the URL
	id := c.Param("id")

	//Get the data off the req
	var body struct {
		Name   string
		Enroll string
		Place  string
		In     string
		Out    string
	}
	c.Bind(&body)

	//Find the post were updating
	var student models.Student
	initializers.DB.First(&student, id)

	//Update it
	initializers.DB.Model(&student).Updates(models.Student{
		Name:   body.Name,
		Enroll: body.Enroll,
		Place:  body.Place,
		In:     body.In,
		Out:    body.Out,
	})

	//repond to them
	c.JSON(200, gin.H{
		"student": student,
	})
}

func StDelete(c *gin.Context) {
	//Get ID off the URL
	id := c.Param("id")

	//Delete the Post
	initializers.DB.Delete(&models.Student{}, id)
	//repond to them
	c.Status(200)
}
