package repository

import (
	"fmt"
	"time"

	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/domain/entity"
	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/db"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db db.DbConfig) *UserRepository {
	return &UserRepository{db.FDB}
}

func (ur *UserRepository) GetExampleUser() (*entity.UserAccount, error) {
	user := entity.UserAccount{
		ID:         1,
		Name:       "John Doe",
		DayOfBirth: "1990-01-01",
		Gender:     "Male",
		Address:    "123 Main St",
		CreatedAt:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	return &user, nil
}

func (ur *UserRepository) RegisterUser(userLoginData *entity.UsersLoginData) (*entity.UsersLoginData, error) {
	if !uniqueEmail(ur.db, userLoginData.EmailAddress) {
		return nil, fmt.Errorf("email already exists")
	}

	if !uniqueLoginName(ur.db, userLoginData.LoginName) {
		return nil, fmt.Errorf("login name already exists")
	}

	err := ur.db.Table("users_login_datas").Create(userLoginData).Error
	if err != nil {
		return nil, err
	}
	return userLoginData, nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*entity.UsersLoginData, error) {
	var user entity.UsersLoginData
	err := ur.db.Table("users_login_datas").Where("email_address = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindUserByID(id int) (*entity.UsersLoginData, error) {
	var user entity.UsersLoginData
	err := ur.db.Table("users_login_datas").Where("user_account_id= ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// TODO: add unique to email and login name in the database
func uniqueEmail(db *gorm.DB, email string) bool {
	var count int64
	db.Table("users_login_datas").Where("email_address = ?", email).Count(&count)
	return count == 0
}

func uniqueLoginName(db *gorm.DB, loginName string) bool {
	var count int64
	db.Table("users_login_datas").Where("login_name = ?", loginName).Count(&count)
	return count == 0
}
