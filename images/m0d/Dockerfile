ARG GO_VER
ARG ALPINE_VER

FROM alpine:${ALPINE_VER} as m0-base
RUN apk add --no-cache tzdata
# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN echo 'hosts: files dns' > /etc/nsswitch.conf

FROM golang:${GO_VER}-alpine${ALPINE_VER} as golang
RUN apk add --no-cache \
	bash \
	binutils-gold \
	gcc \
	g++ \
	git \
	make \
	musl-dev
ADD . $GOPATH/src/github.com/liubaninc/m0
WORKDIR $GOPATH/src/github.com/liubaninc/m0

FROM golang as m0
ARG GO_TAGS
RUN make wasm2c build GO_TAGS=${GO_TAGS}

FROM m0-base
COPY --from=m0 /go/src/github.com/liubaninc/m0/build /usr/local/bin
EXPOSE 26656 26657 26660
RUN apk add --no-cache build-base
CMD ["m0d","start"]