# 错误处理

> `Go` 没有 `try/catch` 异常机制，而是把**错误当作普通的值**来对待（errors are values）。
> 这套"显式返回、显式判断"的风格是 `Go` 工程代码里出现频率最高的一环，也是新手最容易用错的一环。

## `error` 的本质

`error` 就是标准库里一个极简接口：

```go
type error interface {
	Error() string
}
```

- 任何实现了 `Error() string` 的类型都是一个 `error`
- 错误不是语言层面的特殊机制，而是能被**传递、比较、组合**的普通值
- `Go` 的惯用返回签名是 `(T, error)`：`error` 放最后一个返回值，`nil` 表示成功

```go
func parsePositive(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("期望正数，得到 %d", n)
	}
	return n, nil
}
```

- 创建错误：`errors.New("...")` 或 `fmt.Errorf("... %s", x)`
- 编码节奏：**先判 `err`，再用值**——这是最基础也最重要的习惯

## 哨兵错误 (sentinel error)

> 在包级别预先定义、可被判定的错误值，命名以 `Err` 开头。

```go
var ErrNotFound = errors.New("记录不存在")
```

- 标准库里最著名的例子：`io.EOF`、`sql.ErrNoRows`、`os.ErrNotExist`
- 调用方用 `errors.Is(err, ErrNotFound)` 判断"是不是这个特定错误"
- **优先用 `errors.Is` 而不是 `err == ErrNotFound`**：一旦错误被包装，`==` 就失效了，而 `errors.Is` 会沿着错误链逐层比对

## 错误包装与错误链

> 用 `fmt.Errorf` + `%w` 把"底层错误"套进"上层错误"，形成一条**错误链**。

```go
return fmt.Errorf("加载配置 %s 失败: %w", name, err)
```

- `%w` 会建立**可被 `Unwrap` 的链接**；`%v` / `%s` 只是拼文本，链会断掉
- `errors.Is(err, target)`：沿链查找**某个错误值**（哨兵错误），能穿透多层包装
- `errors.As(err, &target)`：沿链查找**某个错误类型**，并把实例取出来
- `errors.Unwrap(err)`：手动剥一层（一般不直接用，理解链结构时有帮助）

> 一句话记忆：**`Is` 比值，`As` 比类型并取值**。

## 自定义错误类型

> 当错误需要携带结构化上下文（字段、错误码等）时，就定义一个类型并实现 `Error()`。

```go
type ValidationError struct {
	Field string
	Value any
	Msg   string
}

func (e *ValidationError) Error() string { ... }
```

- 惯例用**指针接收者**，方便和 `errors.As` 的目标 `**T` 配合
- 想让它挂上错误链，可额外实现 `Unwrap() error`
- 取回结构化信息：`var ve *ValidationError; if errors.As(err, &ve) { ve.Field ... }` —— 比字符串解析可靠得多

## `panic` / `recover`

> **不是常规错误处理手段**。可预期的错误用 `error` 返回；`panic` 只留给"程序进入了不该发生的状态"这类不可恢复的编程错误。

- `recover` **只有在 `defer` 函数里调用才生效**
- 典型用途：在服务/库的**边界**处兜底，防止单个请求的 `panic` 拖垮整个进程
- 结合**命名返回值**把 `panic` 收敛成 `error`：

```go
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("从 panic 中恢复: %v", r)
		}
	}()
	if b == 0 {
		panic(errors.New("除数为 0"))
	}
	return a / b, nil
}
```

## 底层原理解析

- **`error` 是接口，接口是二元组 `(类型, 值)`**：这正是"nil 接口陷阱"的根源（见避坑指南）
- **`errors.New` 返回 `*errorString`**：每次调用都是**不同的指针**，所以哨兵错误必须用**同一个变量**来比较，而不是 `errors.New("x") == errors.New("x")`（永远为 `false`）
- **`%w` 的实现**：`fmt.Errorf` 遇到 `%w` 会返回一个内部的 `*wrapError`，它实现了 `Unwrap() error`。`errors.Is/As` 就是靠不断调用 `Unwrap()` 遍历这条链
- **`errors.Is` 的比对顺序**：先 `==` 比较，再看错误是否实现了 `Is(error) bool` 自定义比较方法，然后 `Unwrap` 进入下一层
- **`panic` 的开销**：`panic` 会展开调用栈(stack unwinding)、执行沿途所有 `defer`，成本远高于返回 `error`，因此不能拿它当控制流

## 避坑指南

- **nil 接口陷阱（最经典）**：返回错误时**返回类型永远写 `error`**，不要写具体的 `*XxxError`；成功路径直接 `return nil`。否则会返回"类型非 nil、值为 nil"的接口，导致调用方 `if err != nil` **恒为真**

```go
var p *MyError = nil
var err error = p
fmt.Println(err == nil) // false！动态类型是 *MyError
```

- **`%w` 写成 `%v`**：错误链会断，`errors.Is/As` 突然失灵。想建立链一定用 `%w`
- **过度包装**：不是每一层都要包。只在**能补充有价值上下文**时才包，否则错误信息会变成一长串重复的废话
- **用字符串匹配判断错误**：`strings.Contains(err.Error(), "not found")` 极其脆弱，文案一改就崩。用哨兵错误 + `errors.Is`
- **哨兵错误暴露内部实现**：导出的 `ErrXxx` 会成为 API 契约的一部分，调用方会依赖它，后续很难改
- **滥用 `panic`**：库/业务逻辑里可预期的失败要用 `error`；`panic` 只用于真正的 bug（数组越界、空指针解引用这类"不该发生"的情况）
- **`recover` 位置错误**：`recover` 不在 `defer` 里、或不在发生 `panic` 的那个 goroutine 里，都无法捕获——**一个 goroutine 的 `panic` 无法被另一个 goroutine `recover`**，会直接崩溃整个进程
