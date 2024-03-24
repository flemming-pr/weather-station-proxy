# Use a base image with the necessary dependencies
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the binary
RUN go build -o weather-station-distributor

# Expose port 3000
EXPOSE 3000

# Set the command to run the binary
CMD ["./weather-station-distributor"]
