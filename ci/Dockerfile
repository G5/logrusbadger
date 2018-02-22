FROM golang:1.10.0

RUN curl -LO https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 && \
    chmod a+x test-reporter-latest-linux-amd64 && \
    mv test-reporter-latest-linux-amd64 /usr/local/bin/cc-test-reporter

RUN mkdir -p /go/src/github.com/g5search/logrusbadger
WORKDIR /go/src/github.com/g5search/logrusbadger

COPY . /go/src/github.com/g5search/logrusbadger
