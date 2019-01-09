FROM golang:1.10-alpine3.7

WORKDIR /go/src/github.com/juliendoutre/go-template

RUN apk add --update make git
RUN go get golang.org/x/tools/cmd/godoc

COPY . .