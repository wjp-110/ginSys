package routes

import (
	"ginSys/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/author/register", controller.Register)
	return r
}
