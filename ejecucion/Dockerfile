FROM golang:1.17

ADD .. /ejecucion

WORKDIR /ejecucion
RUN go mod init cache_test
RUN go mod tidy
RUN go build -o cache_test .
RUN chmod +x /ejecucion

ENTRYPOINT ["/ejecucion/cache_test"]
