package initialize

import (
	"github.com/gin-gonic/gin"
	"go-gin-system/global"
	"go-gin-system/router"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 跨域
	Router.Use(Cors())
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	// 用户表相关的接口
	router.InitUserRouter(ApiGroup)

	global.GVA_LOG.Info("router register success")
	return Router

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id\"")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
