package persistence

import (
	"database/sql"
	"github.com/RioTsukiji/MusicShareBackend/domain"
	"github.com/RioTsukiji/MusicShareBackend/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) InsertUser(DB *sql.DB, name, password string) error {
	stmt, err := DB.Prepare("INSERT INTO user(user_name, password) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, password)
	return err
}

func (up userPersistence) GetByUserName(DB *sql.DB, name string) (*domain.User, error) {
	row := DB.QueryRow("SELECT id, user_name, password, created_at FROM user WHERE user_name = ?", name)
	//row型をgolangで利用できる形にキャストする。
	return convertToUser(row)
}

func convertToUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
