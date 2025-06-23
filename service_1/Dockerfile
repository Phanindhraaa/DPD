# Stage 1: Build the Go app
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

# Force static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Stage 2: Minimal runtime container
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8001

CMD ["./main"]

