FROM golang:1.10.0

ENV DEP_VERSION=v0.4.1

RUN apt-get update && \
    apt-get install -y curl && \
    curl -LO https://github.com/golang/dep/releases/download/$DEP_VERSION/dep-linux-amd64 && \
    chmod a+x dep-linux-amd64 && \
    mv dep-linux-amd64 /usr/local/bin/dep

RUN curl -LO https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 && \
    chmod a+x test-reporter-latest-linux-amd64 && \
    mv test-reporter-latest-linux-amd64 /usr/local/bin/cc-test-reporter

RUN mkdir -p /go/src/github.com/g5search/logrusbadger
WORKDIR /go/src/github.com/g5search/logrusbadger

COPY Gopkg.toml /go/src/github.com/g5search/logrusbadger/Gopkg.toml
COPY Gopkg.lock /go/src/github.com/g5search/logrusbadger/Gopkg.lock

RUN dep ensure -vendor-only

COPY . /go/src/github.com/g5search/logrusbadger
