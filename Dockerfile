FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]