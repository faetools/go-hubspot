FROM golang:{{ .DockerVersion }} as go-builder
RUN apk add --no-cache git
COPY . /go/src/{{ .Name }}
RUN cd /go/src/{{ .Name }}/cmd/app && \
    GOOS=linux CGO_ENABLED=0 go build -o service

FROM ubuntu:20.04
COPY --from=go-builder /go/src/{{ .Name }}/cmd/app/service /usr/local/bin/{{ .Name }}
CMD ["/usr/local/bin/{{ .Name }}"]
