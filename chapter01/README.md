# 基础部分
## 包
  > 包的导入只能单向，否则会产生循环依赖的问题，包的层级尽量简洁

  例如:
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
**int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 `int` 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。**

**`int` 和 `float` 在进行运算时结果会有所不同，应当格外注意**
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

> `var` 关键字和 `:=` 不能同时存在，`:=` 相当于省略了 `var` 关键字并且 `:=` 只能在函数内部使用
### 变量默认值
- 数值类型为 0，
- 布尔类型为 `false`，
- 字符串为 "" (空字符串)

**类型转换(数值类型，需要显示声明(强制))**
~~~go 
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)
  ~~~

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