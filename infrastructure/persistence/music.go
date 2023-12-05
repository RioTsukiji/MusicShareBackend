package persistence

import (
	"database/sql"
	"fmt"
	"github.com/RioTsukiji/MusicShareBackend/domain"
	"github.com/RioTsukiji/MusicShareBackend/domain/repository"
	_ "github.com/go-sql-driver/mysql"
)

type musicPersistence struct{}

func NewMusicPersistence() repository.MusicRepository {
	return &musicPersistence{}
}

func (mp musicPersistence) InsertMusic(DB *sql.DB, title string, artist string, link string, userID int) error {
	stmt, err := DB.Prepare("INSERT INTO songs(song_title, artist_name, song_link, user_id) VALUES(?, ?, ?, ?)")
	fmt.Println(err)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title, artist, link, userID)
	return err
}

// GetAllMusic Todo: FFの音楽だけ取得するようにする
func (mp musicPersistence) GetAllMusic(DB *sql.DB) ([]domain.Music, error) {
	rows, err := DB.Query("SELECT id, song_title, artist_name, song_link, shared_datetime, user_id FROM songs")
	if err != nil {
		return nil, err
	}
	return convertToMusic(rows)
}

func convertToMusic(rows *sql.Rows) ([]domain.Music, error) {
	var allMusic []domain.Music
	for rows.Next() {
		music := domain.Music{}
		err := rows.Scan(&music.ID, &music.Title, &music.Artist, &music.Link, &music.SharedAt, &music.UserID)
		if err != nil {
			return nil, err
		}
		allMusic = append(allMusic, music)
	}
	return allMusic, nil
}
