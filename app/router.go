package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a App) InitRouter() *gin.Engine {

	if !a.Config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     a.Config.App.AllowOrigins,
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		),
	)

	app := router.Group("/v1")
	{
		app.GET("/", a.ExampleHdlr)
	}

	return router
}
