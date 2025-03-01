FROM golang:1.24.0

WORKDIR /app
RUN go install github.com/air-verse/air@latest

COPY  .  .

RUN go mod tidy