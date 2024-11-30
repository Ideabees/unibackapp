package unit

import (
	"errors"
	"otp-app/internal/models"
	"otp-app/internal/repositories"
	"otp-app/internal/services"
	"otp-app/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock UserRepository
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) VerifyUser(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func (m *MockUserRepo) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	service := &services.UserService{Repo: mockRepo}

	user := &models.User{
		MobileNumber: "1234567890",
		Email:        "test@example.com",
		Password:     "securepassword",
	}

	mockRepo.On("CreateUser", user).Return(nil)

	otp, err := service.RegisterUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, otp)
	assert.True(t, utils.ValidateOTPFormat(otp)) // Use a utility to validate OTP format
}

func TestVerifyUserInvalidOTP(t *testing.T) {
	mockRepo := new(MockUserRepo)
	service := &services.UserService{Repo: mockRepo}

	err := service.VerifyUser("test@example.com", "wrong_otp")
	assert.Error(t, err)
	assert.Equal(t, "invalid OTP", err.Error())
}
