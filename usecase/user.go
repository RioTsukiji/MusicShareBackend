package usecase

import (
	"database/sql"
	"github.com/RioTsukiji/MusicShareBackend/domain"
	"github.com/RioTsukiji/MusicShareBackend/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	InsertUser(DB *sql.DB, name string, password string) error
	GetByUserName(DB *sql.DB, name string) (*domain.User, error)
	VerifyPassword(DB *sql.DB, name string, password string) (bool, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) InsertUser(DB *sql.DB, name string, password string) error {

	hashedPassword, hashedErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashedErr != nil {
		return hashedErr
	}

	err := uu.userRepository.InsertUser(DB, name, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (uu userUseCase) GetByUserName(DB *sql.DB, name string) (*domain.User, error) {
	user, err := uu.userRepository.GetByUserName(DB, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) VerifyPassword(DB *sql.DB, name string, password string) (bool, error) {
	user, err := uu.userRepository.GetByUserName(DB, name)
	if err != nil {
		return false, err
	}
	hashedPassword := user.Password
	compareErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if compareErr != nil {
		return false, compareErr
	}
	return true, nil
}
