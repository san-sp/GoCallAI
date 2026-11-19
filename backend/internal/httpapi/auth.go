package httpapi

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/san-sp/GoCallAI/backend/internal/config"
)

func requireAPIToken(cfg config.Config) gin.HandlerFunc {
	expected := []byte(cfg.APIAccessToken)
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		const prefix = "Bearer "
		if !strings.HasPrefix(header, prefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization"})
			return
		}
		provided := []byte(strings.TrimSpace(strings.TrimPrefix(header, prefix)))
		if len(provided) != len(expected) || subtle.ConstantTimeCompare(provided, expected) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization"})
			return
		}
		c.Next()
	}
}
