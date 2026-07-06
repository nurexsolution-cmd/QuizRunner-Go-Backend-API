FROM golang:alpine AS builder

WORKDIR /app
# Copy only the auth service and the shared package it imports.
COPY go.work go.work.sum ./
COPY pkg ./pkg
COPY auth-service ./auth-service

WORKDIR /app/auth-service
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /main .

EXPOSE 8080
CMD ["./main"]
