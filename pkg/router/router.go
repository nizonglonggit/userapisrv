package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 初始化路由
func InitEngine(mode string) *gin.Engine {
	gin.SetMode(mode)
	// 绑定路由，及公共的tomlConfig
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"zero-jwt", "Content-Type", "Origin", "Content-Length"},
		ExposeHeaders:    []string{"zero-jwt"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.Use(gin.Recovery())

	webGroup(engine)

	return engine
}
