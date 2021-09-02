FROM golang:1.17.0-alpine

WORKDIR /app

# auto build & run
RUN go get github.com/githubnemo/CompileDaemon
RUN go get -u github.com/swaggo/swag/cmd/swag
ENTRYPOINT CompileDaemon --build="go build -o /go/bin/main ./server.go" --command="/go/bin/main"