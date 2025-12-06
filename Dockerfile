FROM golang:1.25.5 AS builder

RUN apt-get update && apt-get install -y \
    gcc \
    libc6-dev \
    librdkafka-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

ARG REPOSITORY_NAME

COPY go.mod go.sum ./

COPY monorepo/ ./monorepo/

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o main ${REPOSITORY_NAME:-.}

CMD ["./main"]
