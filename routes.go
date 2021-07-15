package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"xietong.me/ginessential/controller"
	"xietong.me/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	route := r.Group("/api/auth")
	{
		route.POST("/register", controller.Register)
		route.POST("/login", controller.Login)
		route.GET("/info", middleware.AuthMiddleware(), controller.Info)
		route.DELETE("/remove", middleware.AuthMiddleware(), controller.Remove)
	}
	return r
}

/*
 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImV4cCI6MTYyNjk0NzgyNSwiaWF0IjoxNjI2MzQzMDI1LCJpc3MiOiJ4aWV0b25nLm1lIiwic3ViIjoidXNlciB0b2tlbiJ9.a5syFwyh6UHmTB4sprosuJO8o9U63izp4PtL9lgYqsM
*/
