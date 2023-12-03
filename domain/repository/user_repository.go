package repository

import (
	"database/sql"
	"github.com/RioTsukiji/MusicShareBackend/domain"
)

type UserRepository interface {
	InsertUser(DB *sql.DB, name, password string) error
	GetByUserName(DB *sql.DB, name string) (*domain.User, error)
}
