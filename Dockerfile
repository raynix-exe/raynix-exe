FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o http-server-projeto-korp .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/http-server-projeto-korp .
EXPOSE 8080
CMD ["./http-server-projeto-korp"]
