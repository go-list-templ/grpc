FROM golang:1.24.4 as Modules

ADD go.mod go.sum /m/

RUN cd /m && go mod download

FROM golang:1.24.4 as Builder

COPY --from=Modules /go/pkg /go/pkg

RUN useradd -u 10001 go-rest

RUN mkdir -p /go-rest
ADD . /go-rest
WORKDIR /go-rest

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/go-rest ./cmd/go-rest

FROM scratch

COPY --from=Builder /etc/passwd /etc/passwd
USER go-rest

COPY --from=Builder /go-rest/bin/go-rest /go-rest

CMD ["/go-rest"]


