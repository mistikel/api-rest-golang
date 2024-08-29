# Start with the official Golang image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the application on port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]