FROM golang:1.8

COPY ./hello.go /app/

WORKDIR /app

RUN go build hello.go

ENTRYPOINT ./hello
