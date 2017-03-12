FROM golang:1.8
MAINTAINER martiv15@uia.no

RUN mkdir -p /go/src/app

COPY . /go/src/app

WORKDIR /go/src/app

RUN go get github.com/go-martini/martini
RUN go get github.com/martini-contrib/render

RUN go build server.go
RUN go run server.go

EXPOSE 8080
