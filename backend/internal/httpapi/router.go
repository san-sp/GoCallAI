package httpapi

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/san-sp/GoCallAI/backend/internal/config"
)

func NewRouter(_ config.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "gocallai-api",
		})
	})

	router.GET("/readyz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":     "ready",
			"checked_at": time.Now().UTC().Format(time.RFC3339),
			"dependencies": gin.H{
				"database":  "not_configured",
				"redis":     "not_configured",
				"telephony": "not_configured",
				"ai":        "not_configured",
			},
		})
	})

	api := router.Group("/api/v1")
	api.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "GoCallAI",
			"version": "0.1.0",
			"stage":   "foundation",
		})
	})
	return router
}
