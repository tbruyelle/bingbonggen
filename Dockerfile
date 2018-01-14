FROM golang:latest AS builder

RUN go get -u github.com/gopherjs/gopherjs
COPY . /go/src/bingbong

WORKDIR /go/src/bingbong
RUN make

FROM alpine

COPY --from=builder /go/src/bingbong/bingbongd /usr/local/bin
COPY --from=builder /go/src/bingbong/assets /var/bingbong/

EXPOSE 8080

CMD ["/usr/local/bin/bingbongd /var/bingbong"]

