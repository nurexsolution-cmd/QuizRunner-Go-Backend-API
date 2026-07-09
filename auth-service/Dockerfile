FROM golang:alpine AS builder

WORKDIR /app

# 1. Copy workspace files first (cache workspace resolution)
COPY go.work go.work.sum ./

# 2. Copy go.mod/go.sum for ALL services (required by go.work)
COPY auth-service/go.mod auth-service/go.sum ./auth-service/
COPY user-service/go.mod user-service/go.sum ./user-service/
COPY quiz-service/go.mod quiz-service/go.sum ./quiz-service/
COPY result-service/go.mod result-service/go.sum ./result-service/
COPY notification-service/go.mod notification-service/go.sum ./notification-service/
COPY pkg/ ./pkg/

# 3. Download dependencies (cached unless go.mod/go.sum change)
WORKDIR /app/auth-service
RUN go mod download

# 4. Copy the rest of the source code
WORKDIR /app
COPY auth-service/ ./auth-service/

# 5. Build
WORKDIR /app/auth-service
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/auth-service/main .

EXPOSE 8080
CMD ["./main"]
