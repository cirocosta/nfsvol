FROM golang:alpine as builder

ADD ./main /go/src/github.com/cirocosta/nfsvol/main
ADD ./vendor /go/src/github.com/cirocosta/nfsvol/vendor

WORKDIR /go/src/github.com/cirocosta/nfsvol
RUN set -ex && \
  cd ./main && \
  CGO_ENABLED=0 go build -v -a -ldflags '-extldflags "-static"' && \
  mv ./main /usr/bin/nfsvol

FROM busybox
COPY --from=builder /usr/bin/nfsvol /nfsvol

CMD [ "nfsvol" ]
