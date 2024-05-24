package usecase

import (
	"crypto/sha256"
	"fmt"
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = uc.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) UpdateUser(user *model.User) error {
	if user.Password != "" {
		pw := sha256.Sum256([]byte("1234admin"))
		user.Password = fmt.Sprintf("%x", pw)
	} else {
		existingUser, err := uc.userRepository.GetByID(user.ID)
		if err != nil {
			return err
		}
		user.Password = existingUser.Password
	}
	return uc.userRepository.Update(user)
}

func (uc *UserUseCase) GetUserByID(id uint) (*model.User, error) {
	return uc.userRepository.GetByID(id)
}

func (uc *UserUseCase) GetUserByEmail(email string) (*model.User, error) {
	return uc.userRepository.GetByEmail(email)
}

func (uc *UserUseCase) GetAllUsers() ([]*model.User, error) {
	return uc.userRepository.GetAll()
}

func (uc *UserUseCase) PaginateUsers(page int, pageSize int) ([]*model.User, int, error) {
	return uc.userRepository.Paginate(page, pageSize)
}

func (uc *UserUseCase) DeleteUser(user *model.User) error {
	return uc.userRepository.Delete(user)
}

func (uc *UserUseCase) Login(email, password string) (string, error) {
	return uc.userRepository.Login(email, password)
}
