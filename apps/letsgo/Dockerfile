# Start from the official Golang base image in Alpine
FROM golang:alpine AS builder

# Install necessary packages
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy the application source code
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o letsgo

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/letsgo .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the executable
CMD ["/app/letsgo"]
