package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middleware"
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		zap.L().Info("saa:", zap.String("status", settings.Conf.Mode))

		c.String(http.StatusOK, settings.Conf.Version)
	})

	v1 := r.Group("/api/v1")

	v1.POST("/signup", controller.SignUpHandler)

	v1.POST("/login", controller.LoginHandler)

	v1.Use(middleware.JWTAuthMiddleware())

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
	}

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
