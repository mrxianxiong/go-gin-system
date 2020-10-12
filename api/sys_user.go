package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	response "go-gin-system/global"
	"go-gin-system/service"
)

func GetUserList(c *gin.Context) {
	err, list, total := service.GetUserInfoList()
	fmt.Print(list)
	fmt.Print(total)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(nil, c)
	}
}
