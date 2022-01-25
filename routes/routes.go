package routes

import (
	"ginSys/controller"
	"ginSys/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//注册用户
	r.POST("/api/author/register", controller.Register)
	r.POST("/api/author/login", controller.Login)
	r.GET("/api/author/info", middleware.AuthMiddleware(),controller.Info)
	return r
}
