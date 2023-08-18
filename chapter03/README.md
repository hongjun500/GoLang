# 指针

## 指针的定义

- 指针的定义

```go 
var name *Type
```

> `Type` 可以是指针

*此时声明的指针变量还没有分配内存地址，因此没有指向任何变量，所以此时的指针变量的值为 `nil`*

 ~~~go
 // 声明一个指向 int 类型 的指针变量 intPointer
var intPointer *int

// 定义了一个指向 string 类型的指针变量 strPointer
var strPointer = *string
...
~~~

- 指针本质上就是指向某个变量的内存地址的变量
  > 取地址 `&name`
    ~~~go
    str := "hello go"
    将指针 strPointer 指向变量 str 的地址：
    strPointer := &str
    ~~~

  > 取值 `*name`
    ~~~go
    // 最终会打印 ”hello go“
    fmt.Println(*strPointer)
    ~~~
- **指针和变量共享值**，使得在值传递的过程中减少内存消耗（没有值的拷贝）
  > 如果改变了某个变量的值，那么只要指向了这个变量或者是间接指向了这个变量的指针也会随之变量值的改变而改变
- 指针本质上是一个变量，不能指向常量的地址。因为常量的值是不可变的，不能通过指针来改变常量的值

## 内置函数 `new`

- `new` 函数可以创建一个指定类型的指针，语法如下：
    ~~~go
    // 创建一个 int 类型的变量并返回其地址
    var intPointer = new(int)
    // 创建一个 string 类型的变量并返回其地址
    var strPointer = new(string)
    ~~~

- `new` 函数返回的指针指向的值为该类型的零值
    ~~~go
    // 创建一个int类型的变量并返回其地址
    var intPointer = new(int)
    // 创建一个string类型的变量并返回其地址
    var strPointer = new(string)
    // 打印intPointer指向的值
    fmt.Println(*intPointer) // 0
    // 打印strPointer指向的值
    fmt.Println(*strPointer) // ""
    ~~~
- `new` 函数如果指向一个结构体，那么返回的指针指向的结构体的所有字段都是零值
    ~~~go
    // 创建一个结构体并返回其地址
    var structPointer = new(struct {
        name string
        age  int
    })
    // 打印structPointer指向的值
    fmt.Println(*structPointer) // { 0}
    ~~~