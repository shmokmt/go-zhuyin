package zhuyin

import (
	"bytes"

	"github.com/localvar/zhuyin"
	"github.com/mozillazg/go-pinyin"
)

// Convert converts chinese characters to zhuyin（a.k.a bopomofo）
func Convert(s string) string {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone
	chars := pinyin.Convert(s, &a)
	ret := bytes.NewBuffer(make([]byte, 0))
	for i, c := range chars {
		ret.WriteString(zhuyin.PinyinToZhuyin(c[0]))
		if len(chars) == i+1 {
			break
		}
		ret.WriteString(" ")
	}
	return ret.String()
}
