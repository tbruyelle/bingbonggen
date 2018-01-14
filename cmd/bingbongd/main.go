package main

import "bingbong/http"
import "os"

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	panic(http.Listen(path))
}
