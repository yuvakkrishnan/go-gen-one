FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o go-gen-one

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/go-gen-one .
CMD ["./go-gen-one"]
