package main

import (
	"github.com/v1/uniapp/config"
	"github.com/v1/uniapp/internal/handlers"
	"github.com/v1/uniapp/internal/repositories"
	"github.com/v1/uniapp/internal/services"
	"github.com/v1/uniapp/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	userRepo := &repositories.UserRepository{DB: config.DB}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/verify", userHandler.Verify)
	r.POST("/login", userHandler.Login)
	r.POST("/api/uniapp/v1/generate-otp")
	r.POST("/api/uniapp/v1/verify-otp")

	r.GET("/protected", auth.JWTMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the protected route"})
	})

	r.Run(":8080")
}
