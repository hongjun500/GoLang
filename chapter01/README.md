# 基础部分

## 前言

> 查看 `go` 环境

```bash
go env
```

### 解释

- `GOROOT` 安装目录，里面放的就是 `Go` 编译器、标准库源码、工具链等东西
- `GOBIN` 指定 `go install`、`go build -o` 输出的可执行文件路径
  默认：空，`GO` 会把可执行文件放到 `$GOPATH/bin`
  使用场景：
  想把工具直接装进 /usr/local/bin，全局可用。
  在 CI/CD 里避免污染默认 GOPATH/bin。

- `GO111MODULE` 控制 Go Modules 是否启用（Go 1.16 之后默认启用，几乎不用再改，但老项目里常见）。
  - 可选值：
    1. on：强制开启模块模式。
    2. off：强制关闭，走 `GOPATH` 模式。
    3. auto：在有 `go.mod` 文件时开启，否则走 `GOPATH`。

- `GOCACHE` 本地构建缓存目录，用来存放编译过的中间结果, 默认：`$HOME/.cache/go-build`
- `GOPROXY` 设置依赖包下载代理
  1. 默认 https://proxy.golang.org,direct 国外
  2. 改为七牛云提供的 https://goproxy.cn,direct
- `GOSUMDB` Go 的依赖完整性验证中心，保证下载的模块没被篡改,默认：sum.golang.org
- `GOPATH` 以前是 Go 项目的“大本营”，在 `Go Modules` 出现之前，Go 代码必须放在 GOPATH 的子目录里才能编译
  1. 示例结构

  ```bash
  GOPATH/
  bin/      # go install 装的可执行文件
  pkg/      # 编译生成的归档文件（缓存用）
  src/      # 项目的源码（包括依赖）
  ```

  2. Go 1.11 引入 Modules 后，代码可以放在任意位置，不再强制依赖 `GOPATH`
     但是：
     go install 装可执行文件时默认还是会放到 $GOPATH/bin（除非改 GOBIN）。
     有些老项目没切模块化时，还会用到。
  3. 默认值
     Linux/Mac: $HOME/go
     Windows: %USERPROFILE%\go
- `GOMODCACHE` Go modules 依赖的缓存目录
  > go build 或 go get 拉取第三方依赖时，它们会被下载并缓存到这里。
  - 默认路径 `$GOPATH/pkg/mod`

## 包

> 包的导入只能单向，否则会产生循环依赖的问题，包的层级尽量简洁

_错误示范_

```go
    package import_a

    `import (
        "fmt"
        import_b "github.com/hongjun500/GoLang-master/chapter01/import_b"
    )`

    ----------------------------------------------

    package import_b

    import (
      "fmt"
      "github.com/hongjun500/GoLang-master/chapter01/import_a"
    )
    // 上面两个包会产生循环依赖的问题
```

> **大写字母开头的常量、变量在包外可见，小写的则不行**

## 基本数据类型

**int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 `int`
类型，除非你有特殊的理由使用固定大小或无符号的整数类型。**

**`int` 和 `float` 在进行运算时结果会有所不同，应当格外注意**

```go
    // 这两个结果不一样
fmt.Println(1 / 2)
fmt.Println(1 / 2.0)
```

```go
    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // uint8 的别名

    rune // int32 的别名
         // 代表一个 Unicode 码

    float32 float64

    complex64 complex128
```

## 变量

> 变量的申明可以在包级别或函数级别

### `var` 关键字声明

- var 关键字声明可以包含初始值的设定，每个变量一个；若给定了初始值，因为编译器会根据初始值自动推断出类型

### `:=` 短变量声明

> `var` 关键字和 `:=` 不能同时存在，`:=` 相当于省略了 `var` 关键字并且 `:=` 只能在函数内部使用，函数外只能以关键字开头

### 变量默认值

- 数值类型为 0，
- 布尔类型为 `false`，
- 字符串为 "" (空字符串)

**类型转换(数值类型，需要显示声明(强制))**

```go
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)
```

## 常量

- 常量用 `const` 关键字声明
  - `iota` 关键字
    > **`iota` 初始值为 0，每出现一次 `iota` 关键字直到下一个 `const` 关键字出现前 `iota` 都会都会自增 1**
    > **`iota` 可以用来简化定义，在定义枚举时很有用**

## 注释

1. 多行注释 `/* */`
   - 多行注释必须出现在包的顶部
2. 单行注释 `//`

> **单行注释可以出现在任何地方，多行注释必须出现在包的顶部**

- 包的注释简单用 `// path : description` 格式，复杂包的注释可以用 `/* */` 格式
- 函数、方法和结构体的注释用 `// funcName/structName content` 格式
- 变量的注释用 `// varName content` 格式
