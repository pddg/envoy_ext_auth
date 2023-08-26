FROM golang:1.21.0 as builder

WORKDIR /workdir

ENV CGO_ENABLED 0

COPY go.mod /workdir/

RUN go mod download

COPY cmd /workdir/cmd

RUN go build -o hello ./cmd/hello

FROM scratch

COPY --from=builder /workdir/hello /bin/hello

CMD ["/bin/hello"]
