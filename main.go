package main

import (
	"github.com/gin-gonic/gin"
	initializers "github.com/palash0216/Non-Boarders/Initializers"
	"github.com/palash0216/Non-Boarders/controllers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	//Add Controllers
	r.POST("/student", controllers.StCreate)
	r.Run(":3000")
}
