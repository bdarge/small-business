FROM golang:1.16-alpine as dev
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
WORKDIR /bin/app
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o ./bin ./cmd/invoice/" -command="./bin"

FROM golang:1.16-alpine as builder

RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out cmd/invoice/main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /tmp/app/out /app/sm-api

EXPOSE 8080

CMD ["/app/sm-api"]
