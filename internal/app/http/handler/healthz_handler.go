package handler

import (
	"net/http"

	"github.com/FlezzProject/platform-api/internal/infrastructure/iusecase"
	"github.com/gin-gonic/gin"
)

type HealthzHandler struct {
	healthzUsecase iusecase.IHealthzUsecase
}

func NewHealthzHandler(healthzUsecase iusecase.IHealthzUsecase) HealthzHandler {
	return HealthzHandler{
		healthzUsecase: healthzUsecase,
	}
}

func (h HealthzHandler) Show(ctx *gin.Context) {
	status, err := h.healthzUsecase.GetHealthz()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": status})
}
