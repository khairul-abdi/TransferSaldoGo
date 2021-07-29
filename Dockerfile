FROM golang:1.14.0-alpine3.11 AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app/transfersaldo-project

COPY . .

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go build -o binary

CMD ["/app/transfersaldo-project/binary"] 