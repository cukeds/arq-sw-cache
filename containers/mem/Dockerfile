FROM golang:1.17

ADD .. /containers/mem

WORKDIR /containers/mem
RUN go mod init cache_test
RUN go mod tidy
RUN go build -o cache_test .
RUN chmod +x /containers/mem

ENTRYPOINT ["/containers/mem/cache_test"]
