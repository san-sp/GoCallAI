package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/san-sp/GoCallAI/backend/internal/app"
	"github.com/san-sp/GoCallAI/backend/internal/config"
)

func NewRouter(cfg config.Config, application *app.App) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "gocallai-api"})
	})

	router.GET("/readyz", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := application.CheckDependencies(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "not_ready",
				"error":  "one or more dependencies are unavailable",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
			"checked_at": time.Now().UTC().Format(time.RFC3339),
			"dependencies": gin.H{
				"database": "connected",
				"redis": "connected",
				"telephony": "not_configured",
				"ai": "not_configured",
			},
		})
	})

	api := router.Group("/api/v1")
	api.Use(requireAPIToken(cfg))
	api.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "GoCallAI",
			"version": "0.2.0",
			"stage": "backend-persistence",
		})
	})

	settingsHandler := NewSettingsHandler(application.SettingsRepository)
	api.GET("/settings", settingsHandler.Get)
	api.PATCH("/settings", settingsHandler.Update)

	callsHandler := NewCallsHandler(application.CallsRepository)
	api.GET("/calls", callsHandler.List)
	api.GET("/calls/:id", callsHandler.Get)

	return router
}
