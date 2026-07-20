package main

// Reverse 反转字符串。官方 fuzzing 教程里的经典例子：
// 按 rune 反转才能正确处理多字节的 UTF-8 字符（如中文），
// 模糊测试(FuzzReverse)会自动构造刁钻输入来验证"反转两次应还原"这一性质。
func Reverse(s string) string {
	r := []rune(s) // 按 rune 处理，而不是 byte，才不会切坏多字节字符
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
