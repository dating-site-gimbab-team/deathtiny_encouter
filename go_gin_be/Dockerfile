# Step 1: Build the Go application
FROM golang:1.22.6 as builder

WORKDIR /app

# Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o deathtiny_encounters

# Step 2: Create the runtime container
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/deathtiny_encounters .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./deathtiny_encounters"]
