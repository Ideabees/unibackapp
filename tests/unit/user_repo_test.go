package unit

import (
	"database/sql"
	"otp-app/internal/models"
	"otp-app/internal/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := &repositories.UserRepository{DB: db}

	user := &models.User{
		MobileNumber: "1234567890",
		Email:        "test@example.com",
		Password:     "securepassword",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.MobileNumber, user.Email, user.Password, false).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVerifyUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := &repositories.UserRepository{DB: db}

	mock.ExpectExec("UPDATE users SET is_verified = ? WHERE email = ?").
		WithArgs(true, "test@example.com").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.VerifyUser("test@example.com")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
