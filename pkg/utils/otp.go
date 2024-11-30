package utils

import (
	"math/rand"
	"time"
	"fmt"
)

// GenerateOTP creates a random 6-digit OTP
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
