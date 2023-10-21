# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the dependency files
COPY go.mod go.sum ./

# Install project dependencies
RUN go mod download

# Copy the local code to the container
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

# Build the Go application
RUN go build -o bin ./cmd/main/main.go

# Expose a port for your Go application
EXPOSE 8080

# Define the command to run the application
CMD ["./bin"]
