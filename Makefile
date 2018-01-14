all: build

build: gopherjs
	go build -v bingbong/cmd/bingbongd

gopherjs:
	gopherjs build . -o assets/app.js

run: build
	./bingbongd

docker:
	docker build -t bingbong .
