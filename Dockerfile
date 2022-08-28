FROM golang:1.18

WORKDIR /usr/src/app

COPY . .

CMD go run main.go