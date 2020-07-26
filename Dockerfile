FROM golang:1.14-alpine

LABEL maintainer="Rimon Mostafiz <rimonmostafiz@gmail.com>"

RUN apk update && apk add git bash tzdata && rm -rf /var/cache/apk/*
RUN ln -sf /usr/share/zoneinfo/Asia/Dhaka /etc/localtime

WORKDIR /app/flott-bot

COPY . .

RUN go mod download
RUN go build .

CMD ["./flott-bot"]