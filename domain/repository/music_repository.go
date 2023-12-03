package repository

import (
	"database/sql"
	"github.com/RioTsukiji/MusicShareBackend/domain"
)

type MusicRepository interface {
	InsertMusic(DB *sql.DB, title string, artist string, link string, userID int) error
	GetAllMusic(DB *sql.DB) (*domain.Music, error)
}
