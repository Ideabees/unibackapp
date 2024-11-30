package router

import (
	"github.com/v1/uniapp/config"
	"github.com/v1/uniapp/internal/handlers"
	"github.com/v1/uniapp/internal/repositories"
	"github.com/v1/uniapp/internal/services"
	"github.com/v1/uniapp/internal/auth"
	"github.com/gin-gonic/gin"
)




func IntiateRoutes() error {
	config.InitDB()

	userRepo := &repositories.UserRepository{DB: config.DB}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	router := gin.Default()

	router.POST("/register", userHandler.Register)
	router.POST("/verify", userHandler.Verify)
	router.POST("/login", userHandler.Login)
	router.POST("/api/uniapp/v1/generate-otp", userHandler.GenerateOTP)
	router.POST("/api/uniapp/v1/verify-otp", userHandler.VerifyOTP)
	router.POST("/api/uniapp/v1/resend-otp", userHandler.ResendOTP)

	router.GET("/protected", auth.JWTMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the protected route"})
	})

	router.Run(":8080")
	return nil 
}
