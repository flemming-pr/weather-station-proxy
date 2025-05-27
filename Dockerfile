# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files first and download dependencies (caching these layers)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the binary
RUN go build -o weather-station-distributor

# Stage 2: Create a lightweight final image
FROM alpine:3.21

# Add necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set the working directory
WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /app/weather-station-distributor /app/weather-station-distributor

# Expose port 3000
EXPOSE 3000

# Set the command to run the binary
CMD ["./weather-station-distributor"]
