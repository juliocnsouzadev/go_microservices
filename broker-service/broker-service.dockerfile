# base image
FROM golang:1.18-alpine as broker-builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
RUN chmod +x /app/brokerApp

# small image with just executable
FROM alpine:latest
RUN mkdir /app
COPY --from=broker-builder /app/brokerApp /app
CMD ["/app/brokerApp"]