package persistence

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
)

func TestInsertMusic(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO songs").ExpectExec().
		WithArgs("test_title", "test_artist", "test_link", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mp := musicPersistence{}
	mp.InsertMusic(db, "test_title", "test_artist", "test_link", 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAllMusic(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "song_title", "artist_name", "song_link", "shared_datetime", "user_id"}).
		AddRow(1, "test_title", "test_artist", "test_link", time.Now(), 1)

	mock.ExpectQuery("SELECT id, song_title, artist_name, song_link, shared_datetime, user_id FROM songs").
		WillReturnRows(rows)

	mp := musicPersistence{}
	mp.GetAllMusic(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
