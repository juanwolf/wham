FROM golang:1.9.3-stretch

WORKDIR /go/src/wham

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/mitchellh/gox && \
    dep ensure

CMD ["go", "test", "-v", "./cmd"]

