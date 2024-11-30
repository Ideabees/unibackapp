package models

type User struct {
	ID           int    `json:"id"`
	MobileNumber string `json:"mobile_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsVerified   bool   `json:"is_verified"`
}

type GenerateOTP struct {
	MobileNumber string `json:"mobile"`
}

type VerifyOTP struct {
	MobileNumber string `json:"mobile"`
	Otp string `json:"otp"`
}
