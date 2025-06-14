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
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}
