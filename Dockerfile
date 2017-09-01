FROM golang:1.8

ENV GOBIN $GOPATH/bin

RUN  mkdir -p /go/src/github.com/todoer-org/todoer
COPY cmd/ /go/src/github.com/todoer-org/todoer/cmd/
COPY src/ /go/src/github.com/todoer-org/todoer/src/
WORKDIR "/go/src/github.com/todoer-org/todoer/"

RUN go install cmd/reverse/reverse.go
RUN reverse
