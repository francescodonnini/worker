FROM golang:1.19-alpine

RUN apk update

RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN mkdir /pb

COPY ./pb/*.go ./pb/

RUN go build -o /worker

EXPOSE 8080

ENTRYPOINT ["/worker"]
