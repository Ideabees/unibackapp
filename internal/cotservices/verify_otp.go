package cotservices

import (
	"fmt"
	"net/http"
	"io"
)

func VerifyOTP() {

	url := "https://control.msg91.com/api/v5/otp/verify?otp=&mobile="

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authkey", "Enter your MSG91 authkey")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}