FROM golang:1.19 as develop
WORKDIR /go/webtools
COPY ./webtools .
RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]

FROM golang:1.19 as builder

ENV ROOT=/go/webtools
WORKDIR ${ROOT}

COPY ./webtools/go.mod ./webtools/go.sum ./
RUN go mod download

COPY ./webtools ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/webtools

FROM scratch as production
ENV GIN_MODE=release
ENV ROOT=/usr/local/bin
WORKDIR ${ROOT}
COPY --from=builder /go/webtools/webtools ${ROOT}
ENTRYPOINT ["/usr/local/bin/webtools"]