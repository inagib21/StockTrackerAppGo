FROM golang:1.23

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest

COPY . . 

RUN go mod tidy