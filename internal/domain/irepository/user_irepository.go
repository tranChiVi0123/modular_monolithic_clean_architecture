package irepository

import "github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/domain/entity"

type IUserRepository interface {
	GetExampleUser() (*entity.UserAccount, error)
	RegisterUser(user *entity.UsersLoginData) (*entity.UsersLoginData, error)
	FindUserByEmail(email string) (*entity.UsersLoginData, error)
	FindUserByID(id int) (*entity.UsersLoginData, error)
}
