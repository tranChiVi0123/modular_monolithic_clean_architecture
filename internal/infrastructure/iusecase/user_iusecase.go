package iusecase

import (
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/app/presenter"
)

type IUserUsecase interface {
	GetExampleUser() (presenter.UserResponse, error)
	RegisterUserLogin(userLogin presenter.UserRegister, secretKey string) (presenter.AuthResponse, error)
	LoginUser(userLogin presenter.UserLogin, secretKey string) (presenter.AuthResponse, error)
	ValidateUserToken(token string, secretKey string) error
}
