package main

import (
	"errors"
	"fmt"
)

// 错误包装(wrapping)：用 fmt.Errorf + %w 动词把"底层错误"套进"上层错误"里，
// 形成一条错误链(error chain)。这样既能给出高层上下文，又不丢失底层的原始错误。
//
//	上层错误("加载配置失败: %w") → ErrNotFound
//
// 关键点：只有 %w 会建立可被 Unwrap 的链接；%v / %s 只是把文本拼进去，链会断掉。
func wrapError() {
	// 底层：数据层返回哨兵错误
	// 中层：service 加一层上下文并用 %w 包装
	err := loadConfig("app")

	fmt.Println("完整错误信息:", err)
	// 输出类似：加载配置 app 失败: 查询数据库失败: 记录不存在

	// errors.Is：沿着链找有没有某个"目标错误值"。即使 ErrNotFound 被包了两层，依然能命中
	if errors.Is(err, ErrNotFound) {
		fmt.Println("errors.Is 穿透两层包装，仍识别出 ErrNotFound")
	}

	// errors.Unwrap：手动剥一层，一般不直接用，理解链结构时有帮助
	fmt.Println("Unwrap 一层:", errors.Unwrap(err))
}

func loadConfig(name string) error {
	if err := queryDB(name); err != nil {
		// %w 把底层错误挂到链上；文案给出"我这一层在做什么"
		return fmt.Errorf("加载配置 %s 失败: %w", name, err)
	}
	return nil
}

func queryDB(_ string) error {
	// 模拟数据层：包装哨兵错误 ErrNotFound
	return fmt.Errorf("查询数据库失败: %w", ErrNotFound)
}
