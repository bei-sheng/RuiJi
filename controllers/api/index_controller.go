package api

import (
	"github.com/gin-gonic/gin"
	"goRJ/models"
	"strconv"
)

type IndexController struct {
	BaseApiController
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data := models.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	Success(c, maps)
}
func GetUserInfo2(c *gin.Context) {
	setupLogger().Info("这是一个日志")
	request := models.Get("1")
	if request.ID == 0 {
		Failed(c, request)
		return
	}
	Success(c, request)
}
