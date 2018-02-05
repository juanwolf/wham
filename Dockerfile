FROM golang:1.9.3-stretch

WORKDIR /go/src/app

COPY . .

RUN apt-get install -y git && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get github.com/mitchellh/gox && \
    dep ensure

CMD ["go", "test", "-v", "./cmd"]

