FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git
RUN git clone https://github.com/akhtarCareem/watch.git /app
WORKDIR /app
RUN go build -o /request-logger .

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /request-logger /request-logger
EXPOSE 8501
CMD ["/request-logger"]
