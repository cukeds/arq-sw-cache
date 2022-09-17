FROM golang:1.17

ADD .. /containers/loc

WORKDIR /containers/loc
RUN go mod init cache_test
RUN go mod tidy
RUN go build -o cache_test .
RUN chmod +x /containers/loc

ENTRYPOINT ["/containers/loc/cache_test"]
