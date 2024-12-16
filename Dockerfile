# Base image for building the application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Cache module dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fiber-app ./cmd/web

# Final image for running the application
FROM gcr.io/distroless/static:nonroot

# Set the working directory in the final image
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/fiber-app .

# Expose the application port (default Fiber port is 3000)
EXPOSE 3000

# Command to run the application
ENTRYPOINT ["./fiber-app"]
