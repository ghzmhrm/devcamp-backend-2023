FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

RUN go mod download

ENV ENV=Docker

RUN go build -o /Devcamp-2023

EXPOSE 8080

CMD ["/Devcamp-2023"]
