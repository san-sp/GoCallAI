package httpapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/san-sp/GoCallAI/backend/internal/domain"
	"github.com/san-sp/GoCallAI/backend/internal/storage/postgres"
)

type SettingsHandler struct {
	repository *postgres.SettingsRepository
}

func NewSettingsHandler(repository *postgres.SettingsRepository) *SettingsHandler {
	return &SettingsHandler{repository: repository}
}

func (h *SettingsHandler) Get(c *gin.Context) {
	settings, err := h.repository.Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load assistant settings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"settings": settings})
}

func (h *SettingsHandler) Update(c *gin.Context) {
	var input domain.UpdateAssistantSettings
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	if input.Mode != nil {
		switch *input.Mode {
		case "automatic", "manual_active", "disabled":
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "mode must be automatic, manual_active, or disabled"})
			return
		}
	}
	if input.PreferredLanguage != nil {
		switch *input.PreferredLanguage {
		case "auto", "ml", "en":
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "preferred_language must be auto, ml, or en"})
			return
		}
	}
	if input.Greeting != nil {
		*input.Greeting = strings.TrimSpace(*input.Greeting)
		if *input.Greeting == "" || len([]rune(*input.Greeting)) > 240 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "greeting must contain between 1 and 240 characters"})
			return
		}
	}

	settings, err := h.repository.Update(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update assistant settings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"settings": settings})
}
