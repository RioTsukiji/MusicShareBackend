package persistence

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
)

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO users").ExpectExec().
		WithArgs("test_name", "test_pw").
		WillReturnResult(sqlmock.NewResult(1, 1))

	up := userPersistence{}
	if err := up.InsertUser(db, "test_name", "test_pw"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetByUserName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_name", "password", "created_at"}).
		AddRow(1, "testuser", "testpassword", time.Now())

	mock.ExpectQuery("SELECT id, user_name, password, created_at FROM users WHERE user_name =").
		WithArgs("testuser").
		WillReturnRows(rows)

	up := userPersistence{}
	if _, err := up.GetByUserName(db, "testuser"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
