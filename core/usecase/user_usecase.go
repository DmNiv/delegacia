package usecase

import (
	"errors"
	"test/adapter/repository"
	"test/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// CreateUser cria um novo usuário
func (uc *UserUseCase) CreateUser(userRequest *domain.User) (*domain.User, error) {
	existingUser, _ := uc.userRepo.GetByEmail(userRequest.Email)
	if existingUser != nil {
		return nil, errors.New("email já está em uso")
	}

	hashedPassword, err := uc.HashPassword(userRequest.Password)
	if err != nil {
		return nil, errors.New("erro ao gerar hash da senha")
	}

	user := &domain.User{
		FirstName:    userRequest.FirstName,
		LastName:     userRequest.LastName,
		Email:        userRequest.Email,
		Password:     userRequest.Password,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) Login(userRequest *domain.UserLoginDTO) (*domain.User, error) {
	user, err := uc.userRepo.GetByEmail(userRequest.Email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if !uc.CheckPasswordHash(userRequest.Password, user.PasswordHash) {
		return nil, errors.New("senha incorreta")
	}

	return user, nil
}

func (uc *UserUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (uc *UserUseCase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
