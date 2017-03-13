# Uses predefined setup from:
# https://github.com/docker-library/golang/blob/132cd70768e3bc269902e4c7b579203f66dc9f64/1.8/Dockerfile
FROM golang:1.8
MAINTAINER martiv15@uia.no

# Creates root directory of image
RUN mkdir -p /go/src/app

# Copies source into the image
COPY . /go/src/app

# Sets the root for upcoming commands
WORKDIR /go/src/app

# Fetches dependencies
RUN go get github.com/go-martini/martini
RUN go get github.com/martini-contrib/render

# Builds the binary on docker build
RUN go build server.go
ENTRYPOINT ["./server"]
#ENTRYPOINT ["go", "run", "server.go"]

# Exposes port 8080
EXPOSE 8080
