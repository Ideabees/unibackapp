package repositories

import (
	"database/sql"
	"errors"
	"fmt"

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

func (repo *UserRepository) GenerateOTP(mobileNumber string) (string, error) {

	defaultStr := "None"
	defaultFlag := false
	state := "OTP_SEND"

	_, err := repo.DB.Exec(
		"INSERT INTO uniapp.userinfo (mobile_number, email_id, cust_name, company_name, password,  is_mobile_verify, is_email_verify, is_user_verify, mobile_status) VALUES (?,?,?,?,?,?,?,?,?)", mobileNumber, defaultStr, defaultStr, defaultStr, defaultStr, defaultFlag, defaultFlag, defaultFlag, state,
	)

	if err != nil {
		fmt.Println("Error in inserting into db.")
		return "ERROR", err
	}

	fmt.Println("Insertion sucess")

	return "SUCCESS", nil
}
