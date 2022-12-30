# Stage 1
FROM golang:1.20rc1-alpine3.17 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o broker_service ./cmd/api

RUN chmod +x /app/broker_service 

# Stage 2
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker_service /app

CMD ["/app/broker_service"]