FROM golang:alpine AS builder

RUN apk update && \
    apk upgrade && \
    apk add git make

RUN go get -u github.com/gopherjs/gopherjs
COPY . /go/src/bingbong

WORKDIR /go/src/bingbong
RUN make

FROM alpine

# Install ca-certificates for ssl
RUN set -eux; \
	apk add --no-cache --virtual ca-certificates

WORKDIR /app/

COPY --from=builder /go/src/bingbong/bingbongd .
COPY --from=builder /go/src/bingbong/assets ./assets

# not support by heroku
#EXPOSE 8080

RUN adduser -D myuser
USER myuser

CMD ./bingbongd -bind 0.0.0.0:$PORT

