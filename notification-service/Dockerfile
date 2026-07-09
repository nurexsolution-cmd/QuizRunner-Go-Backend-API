FROM golang:alpine AS builder

WORKDIR /app

# 1. Copy workspace files first (cache workspace resolution)
COPY go.work go.work.sum ./

# 2. Copy only go.mod/go.sum for each service + shared pkg — cache dep downloads
COPY notification-service/go.mod notification-service/go.sum ./notification-service/
COPY pkg/ ./pkg/

# 3. Download dependencies (cached unless go.mod/go.sum change)
WORKDIR /app/notification-service
RUN go mod download

# 4. Copy the rest of the source code
WORKDIR /app
COPY notification-service/ ./notification-service/

# 5. Build
WORKDIR /app/notification-service
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/notification-service/main .

EXPOSE 8080
CMD ["./main"]
