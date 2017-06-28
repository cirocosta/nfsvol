FROM golang:alpine as builder

ADD ./ /go/src/github.com/cirocosta/nfsvol
WORKDIR /go/src/github.com/cirocosta/nfsvol
RUN set -ex && \
  CGO_ENABLED=0 go build -v -a -ldflags '-extldflags "-static"' && \
  mv ./nfsvol /usr/bin/nfsvol

FROM busybox
COPY --from=builder /usr/bin/nfsvol /usr/local/bin

CMD [ "nfsvol" ]
