FROM golang:1.19.9

RUN apt-get update && \
    apt-get install bash git

WORKDIR /usr/src/app

EXPOSE 8080

COPY . .
RUN go mod tidy