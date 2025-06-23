package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogogo/pkg/setting"
	"gogogo/routers"
	"net/http"
)

func main() {

	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	server.ListenAndServe()

	//database.InitDB()
	//r := gin.Default()
	////r.POST("/users", controllers.CreateUser)
	////r.GET("/users", controllers.GetUsers)
	////r.GET("/users/:id", controllers.GetUser)
	////r.PUT("/users/:id", controllers.UpdateUser)
	////r.DELETE("/users/:id", controllers.DeleteUser)
	//
	//public := r.Group("/api")
	//{
	//	public.POST("/login", controllers.Login)
	//	public.POST("/register", controllers.Register)
	//}
	//
	//protected := r.Group("/api")
	//protected.Use(middleware.AuthMiddleware())
	//{
	//	r.POST("/users", controllers.CreateUser)
	//	r.GET("/users", controllers.GetUsers)
	//	r.GET("/users/:id", controllers.GetUser)
	//	r.PUT("/users/:id", controllers.UpdateUser)
	//	r.DELETE("/users/:id", controllers.DeleteUser)
	//}
	//
	//r.Run()
}
