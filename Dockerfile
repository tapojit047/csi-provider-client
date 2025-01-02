# Use the official Golang image as the base image
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o csi-provider-client main.go

# Use a minimal image for the runtime
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/csi-provider-client .

# Ensure the binary is executable
RUN chmod +x /app/csi-provider-client

# Expose a Unix socket for the CSI driver communication
CMD ["./csi-provider-client"]
