FROM golang:1.25.7-alpine3.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server

FROM alpine:3.23

WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
