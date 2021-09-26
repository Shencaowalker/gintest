package main

import (
	"gintest1/controller"
	"gintest1/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",controller.Info,middleware.AuthMiddleware(),controller.Info)
	return r
}