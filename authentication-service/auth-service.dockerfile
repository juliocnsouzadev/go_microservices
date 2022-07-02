# base image
FROM golang:1.18-alpine as buider
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o authApp ./cmd/api
RUN chmod +x /app/authApp

# small image with just executable
FROM alpine:latest
RUN mkdir /app
COPY --from=buider /app/authApp /app
CMD ["/app/authApp"]