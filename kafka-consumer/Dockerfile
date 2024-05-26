# Etapa de build
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go mod download
RUN go build -o consumer .

# Etapa final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/consumer .
EXPOSE 8080
CMD ["./consumer"]

