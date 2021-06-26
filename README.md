# go-zhuyin

A utility tool that convert Chinese characters to zhuyin（a.k.a bopomofo）.

## Usage

```go
package main

import (
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
