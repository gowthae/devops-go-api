
# Stage 1: Build the Go application
FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o api .

# Stage 2: Create a minimal runtime image
FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache wget

WORKDIR /app

COPY --from=builder /app/api ./api

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://localhost:8080/health || exit 1

CMD ["./api"]
