# Start from a Go base image
FROM golang:1.24-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules files and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o gitmetricx .

# Expose port 8080 to be accessible outside the container
EXPOSE 8080

# Run the Go application when the container starts
CMD ["go", "run", "main.go"]
