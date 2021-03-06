FROM golang:alpine as builder

ADD ./main /go/src/github.com/cirocosta/nfsvol/main
ADD ./VERSION /go/src/github.com/cirocosta/nfsvol/main/VERSION
ADD ./vendor /go/src/github.com/cirocosta/nfsvol/vendor
ADD ./manager /go/src/github.com/cirocosta/nfsvol/manager

WORKDIR /go/src/github.com/cirocosta/nfsvol
RUN set -ex && \
  cd ./main && \
  CGO_ENABLED=0 go build \
        -tags netgo -v -a \
        -ldflags "-X main.version=$(cat ./VERSION) -extldflags \"-static\"" && \
  mv ./main /usr/bin/nfsvol

FROM busybox
COPY --from=builder /usr/bin/nfsvol /nfsvol

RUN mkdir -p /var/log/nfsvol /mnt/efs

CMD [ "nfsvol" ]
