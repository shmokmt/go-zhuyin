package main

import (
	"flag"
	"fmt"

	"github.com/go-zhuyin"
)

func main() {
	flag.Parse()
	fmt.Println(zhuyin.Convert(flag.Arg(0)))
}
