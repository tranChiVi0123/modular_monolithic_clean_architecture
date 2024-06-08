package iusecase

import (
	"github.com/FlezzProject/platform-api/internal/app/presenter"
)

type IUserUsecase interface {
	GetExampleUser() (presenter.UserResponse, error)
	RegisterUserLogin(userLogin presenter.UserRegister, secretKey string) (presenter.AuthResponse, error)
	LoginUser(userLogin presenter.UserLogin, secretKey string) (presenter.AuthResponse, error)
  ValidateUserToken(token string, secretKey string) error
}
