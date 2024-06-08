package handler

import (
	"net/http"

	"github.com/FlezzProject/platform-api/internal/infrastructure/iusecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsercase iusecase.IUserUsecase
}

func NewUserHandler(userUsecase iusecase.IUserUsecase) UserHandler {
	return UserHandler{
		userUsercase: userUsecase,
	}
}

func (h UserHandler) Show(ctx *gin.Context) {
	user, err := h.userUsercase.GetExampleUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
