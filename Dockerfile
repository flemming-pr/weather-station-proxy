# Use Go 1.23 bookworm as base image
FROM golang:1.24-alpine AS build

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o app

FROM scratch
# Copy binary from the build step
COPY --from=build /build/app /go/bin/app

# Set startup options
EXPOSE 3000
ENTRYPOINT [ "/go/bin/app" ]
