
# Stage 1 - Build
FROM golang:1.23.3 as build

ARG REPOSITORY_NAME

WORKDIR /usr/src/app

COPY ${REPOSITORY_NAME:-.} ./

RUN go mod init main

RUN go build -ldflags "-s -w"

# Stage 2 - Image size reduzed
FROM scratch

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/main /usr/src/app/main

CMD [ "./main"]