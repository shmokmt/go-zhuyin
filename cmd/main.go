package main

import (
	"flag"
	"fmt"

	"github.com/shmokmt/go-zhuyin"
)

func main() {
	flag.Parse()
	fmt.Println(zhuyin.Convert(flag.Arg(0)))
}
