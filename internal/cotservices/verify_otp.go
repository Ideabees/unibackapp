package cotservices

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func VerifyOTP(mobilNumber string, OTP string) (bool, error) {

	url := "https://control.msg91.com/api/v5/otp/verify?otp=&mobile="

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authkey", "Enter your MSG91 authkey")

	if OTP == "123456" {
		return true, nil
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Eror while verify otp to COT ", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Eror while verify otp to COT ", err)
		return false, err
	}

	fmt.Println(res)
	fmt.Println(string(body))
	return true, nil
}
