FROM golang:1.21.0 as builder

WORKDIR /workdir

ENV CGO_ENABLED 0

COPY go.mod /workdir/
COPY go.sum /workdir/

RUN go mod download

COPY cmd /workdir/cmd
COPY pkg /workdir/pkg

RUN go build -o hello ./cmd/hello \
 && go build -o authserver ./cmd/authserver

FROM scratch

COPY --from=builder /workdir/hello /bin/hello
COPY --from=builder /workdir/authserver /bin/authserver

CMD ["/bin/hello"]
