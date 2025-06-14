package main

import (
	"github.com/gin-gonic/gin"
	"gogogo/controllers"
	"gogogo/database"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)

	r.Run()
}
