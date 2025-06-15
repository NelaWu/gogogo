package main

import (
	"github.com/gin-gonic/gin"
	"gogogo/controllers"
	"gogogo/database"
	"gogogo/middleware"
)

func main() {
	database.InitDB()
	r := gin.Default()

	//r.POST("/users", controllers.CreateUser)
	//r.GET("/users", controllers.GetUsers)
	//r.GET("/users/:id", controllers.GetUser)
	//r.PUT("/users/:id", controllers.UpdateUser)
	//r.DELETE("/users/:id", controllers.DeleteUser)

	public := r.Group("/api")
	{
		public.POST("/login", controllers.Login)
		public.POST("/register", controllers.Register)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
	}

	r.Run()
}
