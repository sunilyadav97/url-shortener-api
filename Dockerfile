# Use the official Golang image as the base image.
FROM golang:1.22-alpine

# Install necessary packages
RUN apk add --no-cache git

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the Go application.
RUN go build -o main .

# Expose port 8080.
EXPOSE 8080

# Run the executable.
CMD ["./main"]
