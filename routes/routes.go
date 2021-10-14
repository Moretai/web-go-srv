package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		zap.L().Info("saa:", zap.String("status", settings.Conf.Mode))

		c.String(http.StatusOK, settings.Conf.Version)
	})

	r.POST("/signup", controller.SignUpHandler)

	return r
}
