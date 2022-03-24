FROM golang:1.17-alpine as builder

WORKDIR /tmp/app

COPY go.mod .

COPY go.sum .

RUN go mod download

ADD cmd cmd

RUN ["go", "install", "github.com/swaggo/swag/cmd/swag@latest"]

RUN cd cmd/sm && swag init && cd -

RUN go build -o ./out cmd/sm/main.go

FROM golang:1.17-alpine as dev

WORKDIR /bin/app

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o ./bin/app ./cmd/sm/" -command="./bin/app"

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates git bash

WORKDIR /app

COPY --from=builder /tmp/app/out /app/api

EXPOSE 8080

CMD ["/app/api"]