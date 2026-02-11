FROM golang:1.24.0-alpine AS builder

WORKDIR /app

ENV GOTOOLCHAIN=go1.24.0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/gateway

# =========================

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
