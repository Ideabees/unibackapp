package cotservices

import (
	"fmt"
	"io"
	"net/http"
)

func ResendOTP() {

	url := "https://control.msg91.com/api/v5/otp/retry?authkey=&retrytype=&mobile="

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error in building request")
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error in resend otp")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
