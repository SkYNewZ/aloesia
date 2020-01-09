FROM golang:1.13.5-alpine3.11
WORKDIR /go/src/github.com/skynewz/aloesia

RUN apk add --update --no-cache \
        git \
        make \
        ca-certificates

COPY . .
RUN make build

FROM scratch

ENV PORT 8080

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /go/src/github.com/skynewz/aloesia/bin/aloesia /
ENTRYPOINT ["/aloesia"]