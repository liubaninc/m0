ARG GO_VER
ARG ALPINE_VER
ARG NODE_VER

FROM node:${NODE_VER}-alpine${ALPINE_VER} as m0-base
RUN apk add --no-cache tzdata build-base
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
	musl-dev \
	npm
ADD . $GOPATH/src/github.com/liubaninc/m0
WORKDIR $GOPATH/src/github.com/liubaninc/m0

FROM golang as m0
ARG GO_TAGS
RUN make vue wasm2c build GO_TAGS=${GO_TAGS}

FROM m0-base
COPY --from=m0 /go/src/github.com/liubaninc/m0/build /usr/local/bin
COPY --from=m0 /go/src/github.com/liubaninc/m0/build/wasm /
COPY --from=m0 /go/src/github.com/liubaninc/m0/vue /vue
COPY --from=m0  /go/src/github.com/liubaninc/m0/images/scripts /
COPY --from=m0  /go/src/github.com/liubaninc/m0/images/wait /
# ENV WAIT_VERSION 2.7.2
# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait
EXPOSE 26656 26657 26660 8080 8086 8088
CMD ["m0d","start"]