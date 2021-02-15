FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk update && apk upgrade
RUN apk --no-cache add \
      bash \
      g++ \
      gcc \
      ca-certificates \
      lz4-dev \
      musl-dev \
      cyrus-sasl-dev \
      openssl-dev \
      make \
      python \
      librdkafka-dev \
      pkgconf \
      git
RUN go mod download
RUN go build -tags musl -o main .
CMD ["/app/main"]