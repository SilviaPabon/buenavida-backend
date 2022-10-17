# Golang
FROM golang:alpine3.16

# Watch mode
RUN go install github.com/cosmtrek/air@latest

# Install bash
RUN apk add --no-cache bash

# Configs
ENV GO11MODULE=on \
    GCO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Setup files
RUN mkdir -p /usr/src/buenavida
WORKDIR /usr/src/buenavida
COPY api/ ./api/

WORKDIR /usr/src/buenavida/api
RUN go mod tidy
RUN go mod download
# RUN go build main.go


