package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/iusecase"
)

type AuthMiddleware struct {
	secrectKey   string
	userUsercase iusecase.IUserUsecase
}

func NewAuthMiddleware(userUsecase iusecase.IUserUsecase, secrectKey string) AuthMiddleware {
	return AuthMiddleware{secrectKey, userUsecase}
}

func (am *AuthMiddleware) Execute(ctx *gin.Context) {
	authToken := ctx.GetHeader("auth-token")
	if authToken == "" {
		ctx.JSON(401, fmt.Sprintf("Missing auth token"))
		ctx.Abort()
		return
	}

	err := am.userUsercase.ValidateUserToken(authToken, am.secrectKey)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Invalid token: %s", err))
		ctx.Abort()
		return
	}

	ctx.Next()
}
