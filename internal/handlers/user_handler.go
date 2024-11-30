package handlers

import (
	"net/http"
	"github.com/v1/uniapp/internal/models"
	"github.com/v1/uniapp/internal/services"
	"github.com/v1/uniapp/internal/auth"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func (handler *UserHandler) GenerateOTP (c *gin.Context) {
	var mobileNumber models.GenerateOTP
	if err := c.ShouldBindJSON(&mobileNumber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := handler.Service.GenerateOTP(&mobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

func (handler *UserHandler) ResendOTP (c *gin.Context) {
	var mobileNumber models.GenerateOTP
	if err := c.ShouldBindJSON(&mobileNumber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := handler.Service.GenerateOTP(&mobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

func (handler *UserHandler) VerifyOTP (c *gin.Context) {
	var mobileNumber models.GenerateOTP
	if err := c.ShouldBindJSON(&mobileNumber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := handler.Service.GenerateOTP(&mobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	otp, err := h.Service.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent", "otp": otp}) // OTP shown for testing
}

func (h *UserHandler) Verify(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Service.VerifyUser(req.Email, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User verified"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Service.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, _ := auth.GenerateToken(user)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
