FROM golang:1.15 as builder

WORKDIR /go/src/github.com/vadim-hleif/simple-tcp-server

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install -v -ldflags="-w -s" ./cmd/server

FROM scratch

LABEL \
    name="simple-tcp-sever" \
    maintainer="vadzim.gleif@gmail.com"

ARG WORK_DIR=/simple-tcp-server
WORKDIR $WORK_DIR

ENV SERVER_HOST=0.0.0.0 \
    SERVER_SERVER_PORT=8080
EXPOSE $SERVER_PORT

COPY --from=builder /go/bin/server /bin/server

ENTRYPOINT [ "/bin/server" ]