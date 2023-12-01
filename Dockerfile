FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

RUN go mod download

ENV ENV=Docker

RUN go build -o /devcamp-nsq-2023

EXPOSE 8080

CMD ["/devcamp-nsq-2023"]
