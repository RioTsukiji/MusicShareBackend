package persistence

import (
	"database/sql"
	"github.com/RioTsukiji/MusicShareBackend/domain"
	"github.com/RioTsukiji/MusicShareBackend/domain/repository"
)

type musicPersistence struct{}

func NewMusicPersistence() repository.MusicRepository {
	return &musicPersistence{}
}

func (mp musicPersistence) InsertMusic(DB *sql.DB, title string, artist string, link string, userID int) error {
	stmt, err := DB.Prepare("INSERT INTO songs(song_title, artist_name, song_link, user_id) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title, artist, link, userID)
	return err
}

func (mp musicPersistence) GetAllMusic(DB *sql.DB) (*domain.Music, error) {
	row := DB.QueryRow("SELECT id, song_title, artist_name, song_link, shared_datetime, user_id FROM songs")
	return convertToMusic(row)
}

func convertToMusic(row *sql.Row) (*domain.Music, error) {
	music := domain.Music{}
	err := row.Scan(&music.ID, &music.Title, &music.Artist, &music.Link, &music.SharedAt, &music.UserID)
	if err != nil {
		return nil, err
	}
	return &music, nil
}
