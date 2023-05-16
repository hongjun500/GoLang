package main

import "fmt"

func main() {
	reverse := Reverse("abc")
	fmt.Printf("abc反转得到【%v】\n", reverse)
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
