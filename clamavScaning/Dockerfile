FROM golang:1.18 as base

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod download

RUN go build -o startapiserver ./cmd/apiserver

EXPOSE 8081

CMD ["./startapiserver"]