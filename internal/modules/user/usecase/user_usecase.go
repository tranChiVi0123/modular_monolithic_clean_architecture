package usecase

import (
	"fmt"
	"strconv"

	"github.com/FlezzProject/platform-api/internal/app/presenter"
	"github.com/FlezzProject/platform-api/internal/domain/entity"
	"github.com/FlezzProject/platform-api/internal/domain/irepository"
	repository "github.com/FlezzProject/platform-api/internal/domain/irepository"
	jwt_handler "github.com/FlezzProject/platform-api/pkg/common/jwt"
	"github.com/FlezzProject/platform-api/pkg/common/utils"
)

type UserUsecase struct {
	userRepository irepository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) UserUsecase {
	return UserUsecase{
		userRepository: userRepository,
	}
}

func (u UserUsecase) GetExampleUser() (presenter.UserResponse, error) {
	user, err := u.userRepository.GetExampleUser()
	if err != nil {
		return presenter.UserResponse{}, err
	}

	return presenter.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
	}, nil
}

func (u UserUsecase) RegisterUserLogin(userLogin presenter.UserRegister, secrectKey string) (presenter.AuthResponse, error) {
	passWordHash, err := utils.PasswordBcryptHash(userLogin.Password)
	if err != nil {
		return presenter.AuthResponse{}, err
	}

	userLoginData := entity.UsersLoginData{
		LoginName:                    userLogin.UserName,
		PasswordHash:                 passWordHash,
		EmailAddress:                 userLogin.EmailAddress,
		HasingAlgorithmID:            1,
		EmailAddressVerificationCode: utils.EncodeToString(6),
		EmailVadationStatusID:        3, // NOT_VALIDATED
	}

	user, err := u.userRepository.RegisterUser(&userLoginData)
	if err != nil {
		return presenter.AuthResponse{}, err
	}

	token, err := jwt_handler.CreateJWTToken(user.UserAccountID, secrectKey)
	if err != nil {
		return presenter.AuthResponse{}, err
	}

	return presenter.AuthResponse{Token: token}, nil
}

func (u UserUsecase) LoginUser(userLogin presenter.UserLogin, secrectKey string) (presenter.AuthResponse, error) {
	user, err := u.userRepository.FindUserByEmail(userLogin.EmailAddress)
	if err != nil {
		return presenter.AuthResponse{}, err
	}

	if !utils.PasswordBcryptCompare(userLogin.Password, user.PasswordHash) {
		return presenter.AuthResponse{}, fmt.Errorf("password is incorrect")
	}

	token, err := jwt_handler.CreateJWTToken(user.UserAccountID, secrectKey)
	if err != nil {
		return presenter.AuthResponse{}, err
	}

	return presenter.AuthResponse{Token: token}, nil
}

func (u UserUsecase) ValidateUserToken(token string, secrectKey string) error {
	usersLoginDataIDStr, err := jwt_handler.VerifyJWTToken(token, secrectKey)
	if err != nil {
		return err
	}

	usersLoginDataID, err := strconv.Atoi(usersLoginDataIDStr)
	if err != nil {
		return err
	}

	usersLoginData, err := u.userRepository.FindUserByID(usersLoginDataID)
	fmt.Println(usersLoginData)
	if err != nil {
		return err
	}

	return nil
}
