FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gen-one

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/go-gen-one .
CMD ["./go-gen-one"]
