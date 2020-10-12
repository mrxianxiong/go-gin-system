package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-system/api"
	_ "go-gin-system/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("getUserList", api.GetUserList) // 分页获取用户列表
	}
}
