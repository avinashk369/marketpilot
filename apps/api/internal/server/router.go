package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fintech/marketpilot/apps/api/internal/bootstrap"
)

func setupRouter(container *bootstrap.Container) *gin.Engine {

	router := gin.New()

	router.GET("/api/v1/health", func(c *gin.Context) {

		dbStatus := "UP"
		cacheStatus := "UP"

		if err := container.Database.Health(c.Request.Context()); err != nil {
			dbStatus = "DOWN"
		}

		if err := container.Cache.Health(c.Request.Context()); err != nil {
			cacheStatus = "DOWN"
		}

		c.JSON(http.StatusOK, gin.H{
			"status":      "UP",
			"service":     "marketpilot-api",
			"environment": container.Config.AppEnv,
			"database":    dbStatus,
			"redis":       cacheStatus,
		})
	})

	return router
}