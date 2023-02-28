package routes

import (
	"github.com/gin-gonic/gin"
	"goRJ/controllers/api"
	"goRJ/utils"
)

func InitRouterApi() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // 使用 Logger 中间件、Recovery 中间件
	//r.Use(middlewares.LoggerMiddleware()) //访问中间件日志
	userGroup := r.Group("api/v1")
	{
		userGroup.GET("", api.GetUserInfo2)
	}
	_ = r.Run(utils.HttpPort)
}
