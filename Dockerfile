# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gitmetricx .&#8203;:contentReference[oaicite:2]{index=2}

# :contentReference[oaicite:3]{index=3}
:contentReference[oaicite:4]{index=4}

:contentReference[oaicite:5]{index=5}

# :contentReference[oaicite:6]{index=6}
:contentReference[oaicite:7]{index=7}

# :contentReference[oaicite:8]{index=8}
:contentReference[oaicite:9]{index=9}

# :contentReference[oaicite:10]{index=10}
:contentReference[oaicite:11]{index=11}&#8203;:contentReference[oaicite:12]{index=12}
