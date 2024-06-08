package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/iusecase"
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
