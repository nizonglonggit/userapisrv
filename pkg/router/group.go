package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nizonglonggit/userapisrv/pkg/handler"
)

func webGroup(engine *gin.Engine) {
	v1 := engine.Group("/v1/web/userapisrv")
	{
		v1.POST("/register", handler.UserRegister)
		v1.GET("/user/list", handler.ListAllUser)
	}
}
