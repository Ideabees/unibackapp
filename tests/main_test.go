package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"otp-app/cmd"
	"otp-app/config"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterIntegration(t *testing.T) {
	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()
	cmd.RegisterRoutes(r)

	payload := `{"mobile_number":"1234567890","email":"test@example.com","password":"securepassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "OTP sent")
}

func TestLoginIntegration(t *testing.T) {
	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()
	cmd.RegisterRoutes(r)

	payload := `{"email":"test@example.com","password":"securepassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
