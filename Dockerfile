FROM golang:1.22.5 AS builder

WORKDIR /app

ARG REPOSITORY_NAME

COPY go.mod go.sum ./

COPY ${REPOSITORY_NAME:-.} ${REPOSITORY_NAME:-.}

RUN go mod download

RUN go build -o main ${REPOSITORY_NAME:-.}

CMD ["./main"]
