FROM golang:1.24.0

WORKDIR /src/app

RUN go install github.com/cosmtrek/air@latest

COPY  .  .

RUN go mod tidy