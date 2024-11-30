package cotservices

import (
	"fmt"
	"strings"
	"net/http"
	"io"
	"github.com/v1/uniapp/internal/constant"
)

func SendOTP(mobileNumber string) error {

	baseUrl := constant.MsgBasePath

	endUrl  := fmt.Sprint("?otp_expiry=5&template_id=&mobile=%s&authkey=&realTimeResponse=1", mobileNumber)
	
	finalUrl := baseUrl + endUrl

	//"https://control.msg91.com/api/v5/otp?otp_expiry=5&template_id=&mobile=7892360471&authkey=&realTimeResponse=1"

	payload := strings.NewReader("{\n  \"Param1\": \"value1\",\n  \"Param2\": \"value2\",\n  \"Param3\": \"value3\"\n}")

	req, _ := http.NewRequest("POST", finalUrl, payload)

	req.Header.Add("Content-Type", "application/JSON")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error while making call..")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return nil
}