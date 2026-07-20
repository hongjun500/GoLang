package main

import (
	"errors"
	"fmt"
)

// 哨兵错误(sentinel error)：在包级别预先定义好的、可被 == 比较的错误变量。
// 标准库里最著名的就是 io.EOF、sql.ErrNoRows、os.ErrNotExist。
// 命名约定：以 Err 开头（ErrXxx）。
var (
	ErrNotFound = errors.New("记录不存在")
	ErrPermission = errors.New("没有权限")
)

// findUser 根据 id 返回用户名，找不到时返回哨兵错误 ErrNotFound。
func findUser(id int) (string, error) {
	users := map[int]string{1: "alice", 2: "bob"}
	name, ok := users[id]
	if !ok {
		return "", ErrNotFound
	}
	return name, nil
}

func sentinelError() {
	_, err := findUser(99)

	// 调用方可以用 errors.Is 精确判断"是不是这个特定错误"，从而做差异化处理。
	// 注意：优先用 errors.Is 而不是 err == ErrNotFound，因为一旦错误被包装(见 wrap.go)，
	// 直接用 == 就会失效，而 errors.Is 会沿着错误链逐层比对。
	if errors.Is(err, ErrNotFound) {
		fmt.Println("命中哨兵错误 ErrNotFound，可返回 404 之类的语义")
	}

	// 反例心智模型：不要用字符串匹配来判断错误种类
	// if strings.Contains(err.Error(), "不存在") { ... }  // 脆弱！文案一改就崩
	fmt.Println("原始错误:", err)
}
