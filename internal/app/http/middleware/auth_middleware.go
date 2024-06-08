package middleware

import (
	"fmt"

	"github.com/FlezzProject/platform-api/internal/infrastructure/iusecase"
	"github.com/gin-gonic/gin"
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
