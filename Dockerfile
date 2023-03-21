FROM golang:1.20-alpine3.16 AS builder

RUN mkdir /app
ADD . /app

WORKDIR /app

CMD ["go", "version"]