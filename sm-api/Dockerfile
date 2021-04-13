FROM golang:1.16-alpine as base
WORKDIR /home/app
COPY ./config config

FROM base as dev
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
WORKDIR /bin/app
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o ./bin ./cmd/invoice/" -command="./bin"
