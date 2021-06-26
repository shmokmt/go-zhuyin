# go-zhuyin

[![Go Reference](https://pkg.go.dev/badge/github.com/shmokmt/go-zhuyin.svg)](https://pkg.go.dev/github.com/shmokmt/go-zhuyin)
[![Test](https://github.com/shmokmt/go-zhuyin/actions/workflows/go.yml/badge.svg)](https://github.com/shmokmt/go-zhuyin/actions/workflows/go.yml)

A utility tool that convert Chinese characters to zhuyin（a.k.a bopomofo）.

## Usage

```go
package main

import (
	"fmt"

	"github.com/shmokmt/go-zhuyin"
)

func main() {
	fmt.Print(zhuyin.Convert("日本人")) // ㄖˋ ㄅㄣˇ ㄖㄣˊ
}
```


```
$ zhuyin 日本人
ㄖˋ ㄅㄣˇ ㄖㄣˊ
```

## Author

shmokmt
