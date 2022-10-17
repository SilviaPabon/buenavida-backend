# Golang compiler
FROM golang:alpine3.16 AS builder

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
RUN go build main.go

# Golang runner
FROM alpine:3.16 AS runner
RUN apk add --no-cache bash

WORKDIR /apt/api
COPY --from=builder /usr/src/buenavida/api .
RUN ./main
