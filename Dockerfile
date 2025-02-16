# Stage 1: Build and test the application
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Run tests
RUN go test -v ./...

# Build the application
RUN go build -o link-in-bio-api .

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/link-in-bio-api .

# no cred. etc so putting env in container
COPY .env .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./link-in-bio-api"]