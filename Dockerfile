FROM golang:latest

WORKDIR /go/src/github.com/ceiba-meli-demo/movies

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 8081