# Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app

# Copy only go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o go-gen-one

# Stage 2: Minimal runtime image
FROM alpine:3.18

# Install ca-certificates if your app makes HTTPS calls
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/go-gen-one .

# Command to run
CMD ["./go-gen-one"]
