FROM golang:1.17-alpine

RUN go version 
ENV GO_PATH=./

COPY ./ ./

RUN go mod download

RUN go build -o /app
