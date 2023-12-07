package main

import (
	initializers "github.com/palash0216/Non-Boarders/Initializers"
	"github.com/palash0216/Non-Boarders/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Student{})
}
