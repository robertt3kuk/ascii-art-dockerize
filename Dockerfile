# syntax=docker/dockerfile:1
FROM golang:alpine AS builder 
WORKDIR /app
COPY . .
RUN go mod download
RUN go build  -o server
FROM alpine
LABEL version="0.1"
LABEL maintaner="robertt3kuk "
WORKDIR /app
COPY --from=builder /app .
EXPOSE 3000
CMD [ "./server" ]
