FROM golang:1.8

ENV GOBIN $GOPATH/bin

RUN  mkdir -p /go/src/github.com/todoer-org/todoer/tests/integration/api
COPY . /go/src/github.com/todoer-org/todoer/tests/integration/api
WORKDIR "/go/src/github.com/todoer-org/todoer/tests/integration/api"