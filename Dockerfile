# Use the official Golang image as a base
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o gin-gonic-rest main.go

# Use base image for the final container
FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat

# Set the working directory for the final image
WORKDIR /app

# Copy the built binary
COPY --from=builder /app/gin-gonic-rest /app/gin-gonic-rest

# Ensure the binary has execute permissions
RUN chmod +x /app/gin-gonic-rest

# Expose the port
EXPOSE 8765

# Command to run the binary
CMD ["/app/gin-gonic-rest"]
