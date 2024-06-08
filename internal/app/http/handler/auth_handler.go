package handler

import (
	"net/http"

	"github.com/FlezzProject/platform-api/internal/app/presenter"
	"github.com/FlezzProject/platform-api/internal/infrastructure/iusecase"
	errors_handler "github.com/FlezzProject/platform-api/pkg/common/errors"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	secretKey   string
	userUsecase iusecase.IUserUsecase
}

func NewAuthHandler(userUsecase iusecase.IUserUsecase, secretKey string) AuthHandler {
	return AuthHandler{
		secretKey:   secretKey,
		userUsecase: userUsecase,
	}
}

func (h AuthHandler) Register(ctx *gin.Context) {
	var userRegister presenter.UserRegister

	if err := ctx.Bind(&userRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, errors_handler.New400ErrorResponse(err))
		return
	}

	authResponse, err := h.userUsecase.RegisterUserLogin(userRegister, h.secretKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors_handler.New400ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, presenter.AuthResponse{
		Token: authResponse.Token,
	})
}

func (h AuthHandler) Login(ctx *gin.Context) {
	var userLogin presenter.UserLogin

	if err := ctx.Bind(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, errors_handler.New400ErrorResponse(err))
		return
	}

	authResponse, err := h.userUsecase.LoginUser(userLogin, h.secretKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors_handler.New400ErrorResponse(err))
		return
	}

  ctx.JSON(http.StatusOK, presenter.AuthResponse{
    Token: authResponse.Token,
  })
}

func (h AuthHandler) Logout(ctx *gin.Context) {
	// code here
}
