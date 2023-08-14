# 流程控制语句

## `if/else`
```go
if x < y {
    return x
} else {
    return y
}
```

- `if` 语句的条件表达式不需要使用小括号包裹，但是大括号必须有;  
- `if` 语句的条件表达式可以声明一个短变量，该变量可以在 `if` 结构体及其紧跟其后的
`else` 结构体中使用(即为遇到了 `return` 后就无法使用)

## `for` 循环
```go
// 普通标准循环
for initialization; condition; post {
    // zero or more statements
}

// 省略初始化表达式和后置语句
i := 100 // 本质上这里的 i 赋值 100 就是初始化语句
for i > 0 {
	i-- // 而这里的 i-- 则是后置语句，作用其实就是改变初始语句赋的值
    // 使其条件表达式的值为 false 结束循环，否则将陷入无限循环
}
```

   - 目前循环只有 for (v1.19)
   - `for` 循环结构体包含初始化表达式，条件表达式，后置语句;
   用英文分号 `;` 隔开;其中只有条件表达式是必须包含的
   - 后置语句也应该显示声明(比如使用 `break` 结束所在循环，`continue` 跳出循环，或者使用 `return` 退出循环)，否则容县陷入无限循环
  > `for` 语句同 `if` 语句一样可以没有条件表达式的小括号，大括号必须

### `range` 关键字遍历循环
```go
for k, v range coll { 
	// ...
}
```
- **`range` 关键字遍历循环可以遍历数组，切片，字符串，map 及通道 `channel`**
- **`range` 关键字遍历循环返回两个值，第一个是索引，第二个是索引对应的值**
~~`range` 关键字遍历循环可以使用 `_` 忽略索引或者值~~
- `range` 关键字遍历循环可以使用 `_` 忽略索引
- `range` **遍历循环时如果只取索引，可以直接省略第二个返回值**
- `range` 关键字遍历循环可以使用 `break` 结束循环，使用 `continue` 跳出循环，使用 `return` 退出循环

## `switch` 语句
```go
switch {
    case num < 0:
        fmt.Println("Number is negative")
    case num == 0:
        fmt.Println("Number is zero")
    case num > 0:
        fmt.Println("Number is positive")
}
```
- 可以没有表达式，没有表达式的 `switch` 语句相当于 `if/else` 语句
- 表达式可以是任何类型，不限于常量或者整数
- 从上向下执行，直到找到匹配项，此时 `switch` 程序运行结束
- 可以不包含 `default` 关键字语句
- 每个 case 后面会默认包含一个 `break` 关键字语句
- 如果存在 `fallthrough` 语句，那么其后必须包含一个 `case` 或者是 `default` ，否则编译不通过程序报错
  > `fallthrough` 关键字语句若是出现在符合表达式结果的 `case` 结构体中,那么就能够发挥出作用：
  > 立即执行下一个 `case` 结构体代码（如果有 `case`），不管这个 `case` 能否命中表达式结果都必须执行
  > `fallthrough` 后面必须有一个 `case`（最少一个）或者是直接跟一个 `default` 关键字(可以没有代码，但是关键字必须有)
   
## `defer` 语句 
> 延迟函数：*基于函数而非代码块*
- defer 语句会将函数推迟到外层函数返回之后执行 
- 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用 
- 多个 `defer` 推迟的函数调用会被**压入一个栈**中。当`外层函数`返回时，被推迟的函数会按照**`**后进先出**`**的顺序调用

*例子*

```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
	// 执行步骤 6 将 un 函数压入栈中，栈2
	// 执行步骤 7 执行函数 trace("a")，输出 entering: a 并将返回值 "a" 传递给 un 函数
    defer un(trace("a"))
	// 执行步骤 8
    fmt.Println("in a")
}

func b() {
	// 执行步骤 2 将 un 函数压入栈中，栈1
	// 执行步骤 3 执行函数 trace("b")，输出 entering: b 并将返回值 "b" 传递给 un 函数
    defer un(trace("b"))
	// 执行步骤 4
    fmt.Println("in b")
	// 执行步骤 5
    a()
}

func main() {
	// 执行步骤 1
    b()
}
```
**最终结果**
```go
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```
*对于上方的 un 函数执行顺序：un("b")的顺序先于 un("a"), 因此 `leaving: a` 先输出*

## `goto` 语句
```go
func gotoFunc() {
    i := 0
HERE:
    fmt.Println(i) // i = 0 打印出0
    i++            // 此时 i = 1
    if i < 10 {    // i = 1会进入此判断结构体
        // 跳到HERE	标签处，执行打印，程序会接着执行i++,又进入条件判断以此类推直到i=10的时候程序运行结束
        goto HERE
    }
}
```
- `goto` 语句可以无条件地转移到**同一函数**中的带有**自定义标签**的语句处执行
- `goto` 语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能