# Use the official Golang image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o unibackapp

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./unibackapp"]
