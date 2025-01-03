# unibackapp
It contains backend code

Curl for generate-otp

curl --location 'http://localhost:8080/api/uniapp/v1/generate-otp' \
--header 'Content-Type: application/json' \
--data '{
    "mobile_number": "3434322"
}'

curl --location 'http://localhost:8080/api/uniapp/v1/resend-otp' \
--header 'Content-Type: application/json' \
--data '{
    "mobile_number": "3434322"
}'

Run Docker Compose file
docker-compose -f docker-compose.uniapp.yml up --build



