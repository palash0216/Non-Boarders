package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/palash0216/Non-Boarders/controllers"
	"github.com/stretchr/testify/assert"
)

func TestStCreate(t *testing.T) {
	// Initialize a test router
	router := gin.New()
	router.POST("/student", controllers.StCreate)

	// Create a test request with a sample JSON payload for success
	successJSON := `{
	"Name": "Rahul Sharma",
	"Enroll": "102V202",
	"Place": "Raogarh",
	"In": "8:55 AM",
	"Out": "4:55 PM"}`
	reqSuccess, err := http.NewRequest("POST", "/student", bytes.NewBufferString(successJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request
	wSuccess := httptest.NewRecorder()
	router.ServeHTTP(wSuccess, reqSuccess)

	// Assert the HTTP status code for success
	assert.Equal(t, http.StatusOK, wSuccess.Code)

	// Assert the response body for success
	// Add your specific assertions here based on the expected response for success

	// Create a test request with a sample JSON payload for an error
	errorJSON := `{
	"Name": "Rahul Sharma",
	"Enroll": "102V202",
	"Place": "Raogarh",
	"In": "8:55 AM",
	"Out": "4:55 PM"}`
	reqError, err := http.NewRequest("POST", "/student", bytes.NewBufferString(errorJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request
	wError := httptest.NewRecorder()
	router.ServeHTTP(wError, reqError)

	// Assert the HTTP status code for error
	assert.Equal(t, http.StatusBadRequest, wError.Code)

	// Assert the response body for error
	// Add your specific assertions here based on the expected response for error
}

// func TestStudentIndex(t *testing.T) {
// 	// Initialize a test router
// 	router := gin.New()
// 	router.GET("/student", controllers.StudentIndex)

// 	// Create a test request
// 	req, err := http.NewRequest("GET", "/student", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Perform the request
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assert the HTTP status code
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Add more assertions as needed based on your specific logic
// }

// Similarly, create tests for StudentShow, StUpdate, and StDelete functions...
