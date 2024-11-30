package repositories

import (
	"database/sql"
	"errors"
	"github.com/v1/uniapp/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	_, err := repo.DB.Exec(
		"INSERT INTO users (mobile_number, email, password, is_verified) VALUES (?, ?, ?, ?)",
		user.MobileNumber, user.Email, user.Password, false,
	)
	return err
}

func (repo *UserRepository) VerifyUser(email string) error {
	res, err := repo.DB.Exec("UPDATE users SET is_verified = ? WHERE email = ?", true, email)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (repo *UserRepository) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.QueryRow(
		"SELECT id, is_verified FROM users WHERE email = ? AND password = ?",
		email, password,
	).Scan(&user.ID, &user.IsVerified)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GenerateOTP(mobileNumber string, flag bool) (string, error) {
	
	/*err := repo.DB.QueryRow(
		"Insert into USERINFO id, is_verified FROM users WHERE email = ? AND password = ?",
		email, password,
	).Scan(&user.ID, &user.IsVerified)
	if err != nil {
		return nil, err
	}*/

	return "SUCCESS", nil
}