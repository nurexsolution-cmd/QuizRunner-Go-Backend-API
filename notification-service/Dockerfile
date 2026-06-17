FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.work go.work.sum* ./
COPY pkg/ pkg/
COPY notification-service/ notification-service/

WORKDIR /app/notification-service
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/notification-service/main .

EXPOSE 8080
CMD ["./main"]
