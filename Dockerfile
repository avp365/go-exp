#---Build stage---
FROM golang:1.19.3 AS builder
WORKDIR /app
COPY . /app
RUN go build -o /go/bin/app ./cmd/app/main.go
ENTRYPOINT ["/go/bin/app"] 