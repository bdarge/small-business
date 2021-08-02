FROM golang:1.16-alpine as builder

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out cmd/invoice/main.go

FROM golang:1.16-alpine as dev

WORKDIR /bin/app

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o ./bin ./cmd/invoice/" -command="./bin"

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY --from=builder /tmp/app/out /app/sm-api

EXPOSE 8080

CMD ["/app/sm-api"]
