FROM golang:1.23-alpine AS builder

WORKDIR /app
# We need to copy the workspace files and the shared pkg directory
COPY go.work go.work.sum* ./
COPY pkg/ pkg/
COPY auth-service/ auth-service/

WORKDIR /app/auth-service
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/auth-service/main .

EXPOSE 8080
CMD ["./main"]
