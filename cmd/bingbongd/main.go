package main

import (
	"bingbong/http"
	"flag"
)

var (
	bind = flag.String("bind", ":8080", "The http port")
	res  = flag.String("res", "", "The resources path")
)

func main() {
	flag.Parse()
	panic(http.Listen(*bind, *res))
}
