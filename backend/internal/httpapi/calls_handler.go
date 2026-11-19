package httpapi

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/san-sp/GoCallAI/backend/internal/storage/postgres"
)

type CallsHandler struct {
	repository *postgres.CallsRepository
}

func NewCallsHandler(repository *postgres.CallsRepository) *CallsHandler {
	return &CallsHandler{repository: repository}
}

func (h *CallsHandler) List(c *gin.Context) {
	limit := 20
	offset := 0

	if raw := c.Query("limit"); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil || value < 1 || value > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be between 1 and 100"})
			return
		}
		limit = value
	}
	if raw := c.Query("offset"); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil || value < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "offset must be zero or greater"})
			return
		}
		offset = value
	}

	calls, err := h.repository.List(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load call history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"calls":  calls,
		"limit":  limit,
		"offset": offset,
	})
}

func (h *CallsHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid call ID"})
		return
	}

	call, err := h.repository.Get(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "call not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load call"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"call": call})
}
