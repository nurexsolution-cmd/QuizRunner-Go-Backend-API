FROM golang:alpine AS builder

WORKDIR /app

# 1. Copy workspace files first (cache workspace resolution)
COPY go.work go.work.sum ./

# 2. Copy only go.mod/go.sum for each service + shared pkg — cache dep downloads
COPY quiz-service/go.mod quiz-service/go.sum ./quiz-service/
COPY pkg/ ./pkg/

# 3. Download dependencies (cached unless go.mod/go.sum change)
WORKDIR /app/quiz-service
RUN go mod download

# 4. Copy the rest of the source code
WORKDIR /app
COPY quiz-service/ ./quiz-service/

# 5. Build
WORKDIR /app/quiz-service
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/quiz-service/main .

EXPOSE 8080
CMD ["./main"]
