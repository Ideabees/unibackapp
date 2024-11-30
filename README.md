# unibackapp
It contains backend code

Curl for generate-otp

curl --location 'http://localhost:8080/api/uniapp/v1/generate-otp' \
--header 'Content-Type: application/json' \
--data '{
    "mobile_number": "3434322"
}'