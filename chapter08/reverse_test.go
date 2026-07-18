package main

import (
	"testing"
	"unicode/utf8"
)

/*
单元测试
*/
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"hello, world", "dlrow , olleH"},
		{"", ""},
		{"123", "789"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

/*
模糊测试
*/
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func Reverse(s string) string {
	// 将字符串 s 转换成 byte 数组
	b := []byte(s)
	/*for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}*/
	// 双指针分别指向首尾
	i := 0
	j := len(b) - 1
	for i < len(b)/2 {
		// 交换位置
		b[i], b[j] = b[j], b[i]
		i = i + 1
		j = j - 1
	}
	return string(b)
}
