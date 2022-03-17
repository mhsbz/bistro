package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("main...")
	fasthttp.Get(nil, "")
}
