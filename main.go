package main

import (
	"net/url"
)

var Cnt = 0

func main() {
	parsedUrl, err := url.Parse("xxxxxxx:8888")
	print(err)
	print("\n")
	print(parsedUrl.Scheme)
	print("\n")
	print(parsedUrl.Host)
	print("\n")
	print(parsedUrl.Port())
	print("\n")
	print(parsedUrl.Path)
	print("\n")
}
