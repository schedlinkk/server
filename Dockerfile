# Use a lightweight official Go image as the build environment
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files to download dependencies first (for caching)
# COPY go.sum ./
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go application from the ./cmd/server directory
# This command creates an executable named "server"
RUN go build -o server ./cmd/server

# Use a minimal base image for the final container
FROM alpine:latest

# Set a working directory in the final image
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose port 8080 so that the application is accessible
EXPOSE 8080

# Run the Go server when the container starts
CMD ["./server"]
