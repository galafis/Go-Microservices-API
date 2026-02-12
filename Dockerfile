# Build stage
FROM golang:1.22-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/Go-Microservices-API ./...

# Runtime stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/Go-Microservices-API /app/Go-Microservices-API

EXPOSE 8080

CMD ["/app/Go-Microservices-API"]
