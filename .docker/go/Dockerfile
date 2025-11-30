FROM golang:1.24.4 as Modules

ADD go.mod go.sum /m/

RUN cd /m && go mod download

FROM golang:1.24.4 as Builder

COPY --from=Modules /go/pkg /go/pkg

RUN useradd -u 10001 go-http-servere

RUN mkdir -p /go-http-servere
ADD . /go-http-servere
WORKDIR /go-http-servere

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/go-http-servere ./cmd/go-http-servere

FROM scratch

COPY --from=Builder /etc/passwd /etc/passwd
USER go-http-servere

COPY --from=Builder /go-http-servere/bin/go-http-servere /go-http-servere

CMD ["/go-http-servere"]


