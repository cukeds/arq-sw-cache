FROM golang:1.17

ADD . /containers

WORKDIR /containers
RUN go mod init cache_test
RUN go mod tidy
RUN go build -o cache_test .
RUN chmod +x /containers

ENTRYPOINT ["/containers/cache_test"]
