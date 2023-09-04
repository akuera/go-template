FROM golang:1.18.0-alpine as builder

RUN apk update -q && apk add -q git curl automake libtool make g++ unzip wget

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENNABLED=0

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -tags release -o bin/application main.go

FROM alpine

RUN apk add --update \
    && apk add curl

COPY --from=builder /app/bin/application /bin/application

EXPOSE 5000
CMD ["/bin/application"]
