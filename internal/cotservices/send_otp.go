package cotservices

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func SendOTP() {

	url := "https://control.msg91.com/api/v5/otp?otp_expiry=5&template_id=&mobile=7892360471&authkey=&realTimeResponse=1"

	payload := strings.NewReader("{\n  \"Param1\": \"value1\",\n  \"Param2\": \"value2\",\n  \"Param3\": \"value3\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/JSON")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}