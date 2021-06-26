package zhuyin

import (
	"testing"
)

func TestConvert(t *testing.T) {
	want := "ㄖˋ ㄅㄣˇ ㄖㄣˊ"
	got := Convert("日本人")
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
