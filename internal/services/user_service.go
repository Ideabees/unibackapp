package services

import (
	"errors"
	"fmt"

	"github.com/v1/uniapp/internal/cotservices"
	"github.com/v1/uniapp/internal/models"
	"github.com/v1/uniapp/internal/repositories"
	"github.com/v1/uniapp/pkg/utils"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) GenerateOTP(user *models.GenerateOTP) (string, error) {

	// call cot api
	err := cotservices.SendOTP(user.MobileNumber)
	if err != nil {
		fmt.Println("Error in send otp service layer")
		return "Error", err
	}
	// call repo layer to update the status

	msg, err := s.Repo.GenerateOTP(user.MobileNumber)

	if err != nil {
		fmt.Println("Error while insering into db")
		return msg, err
	}

	return msg, nil
}

func (s *UserService) VerifyOTP(user *models.User) (string, error) {
	if err := s.Repo.CreateUser(user); err != nil {
		return "", err
	}
	otp := utils.GenerateOTP()
	// In production, send OTP via SMS or Email
	return otp, nil
}

func (s *UserService) RegisterUser(user *models.User) (string, error) {
	if err := s.Repo.CreateUser(user); err != nil {
		return "", err
	}
	otp := utils.GenerateOTP()
	// In production, send OTP via SMS or Email
	return otp, nil
}

func (s *UserService) VerifyUser(email, otp string) error {
	if otp != "123456" { // Mock OTP validation
		return errors.New("invalid OTP")
	}
	return s.Repo.VerifyUser(email)
}

func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	return s.Repo.GetUserByEmailAndPassword(email, password)
}
