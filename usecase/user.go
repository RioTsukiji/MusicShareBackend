package usecase

import (
	"database/sql"
	"github.com/RioTsukiji/MusicShareBackend/domain"
	"github.com/RioTsukiji/MusicShareBackend/domain/repository"
)

type UserUseCase interface {
	InsertUser(DB *sql.DB, name string, password string) error
	GetByUserName(DB *sql.DB, name string) (*domain.User, error)
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
	//一意でランダムな文字列を生成する
	hashedPassword, err := uuid.NewRandom() //返り値はuuid型
	if err != nil {
		return err
	}

	//domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err = uu.userRepository.InsertUser(DB, name, password)
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
