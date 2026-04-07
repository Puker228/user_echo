package handler

import (
	"net/http"

	"github.com/Puker228/user_echo/internal/domain"
	"github.com/Puker228/user_echo/internal/usecase"
	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	uc usecase.StatsUseCase
}

func NewStatsHandler(uc usecase.StatsUseCase) *StatsHandler {
	return &StatsHandler{uc: uc}
}

func (h *StatsHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/stats", h.save)
}

func (h *StatsHandler) save(c *gin.Context) {
	var stats domain.UserStats
	if err := c.ShouldBindJSON(&stats); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Save(c.Request.Context(), stats); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "saved"})
}
