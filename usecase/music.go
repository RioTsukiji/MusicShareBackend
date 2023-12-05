package usecase

import (
	"database/sql"
	"github.com/RioTsukiji/MusicShareBackend/domain"
	"github.com/RioTsukiji/MusicShareBackend/domain/repository"
)

type MusicUseCase interface {
	InsertMusic(DB *sql.DB, title string, artist string, link string, userID int) error
	GetAllMusic(DB *sql.DB) ([]domain.Music, error)
}

type musicUseCase struct {
	musicRepository repository.MusicRepository
}

func NewMusicUseCase(mr repository.MusicRepository) MusicUseCase {
	return &musicUseCase{
		musicRepository: mr,
	}
}

func (mu musicUseCase) InsertMusic(DB *sql.DB, title string, artist string, link string, userID int) error {
	//本来ならemailのバリデーションをする

	//domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err := mu.musicRepository.InsertMusic(DB, title, artist, link, userID)
	if err != nil {
		return err
	}
	return nil
}

func (mu musicUseCase) GetAllMusic(DB *sql.DB) ([]domain.Music, error) {
	music, err := mu.musicRepository.GetAllMusic(DB)
	if err != nil {
		return nil, err
	}
	return music, nil
}
