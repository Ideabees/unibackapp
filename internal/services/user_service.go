package services

import (
	"errors"
	"github.com/v1/uniapp/internal/models"
	"github.com/v1/uniapp/internal/repositories"
	"github.com/v1/uniapp/pkg/utils"
)

type UserService struct {
	Repo *repositories.UserRepository
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
