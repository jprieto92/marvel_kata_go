# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o marvel-kata-k8s cmd/k8s/main.go
## Our start command which kicks off
## our newly created binary executable
## EXPOSE 8080

CMD ["/app/marvel-kata-k8s"]