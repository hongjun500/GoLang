### 第一章节
- 包
- 变量
- var关键字, := 短变量声明
   > var关键字和:=不能同时存在，:=相当于省略了var关键字并且:=只能在函数内部使用 
  - 基本数据类型

      ~~~shell
       bool
    
      string
    
      int  int8  int16  int32  int64
      uint uint8 uint16 uint32 uint64 uintptr(指针)
    
      byte // uint8 的别名
    
      rune // int32 的别名
          // 表示一个 Unicode 码点
    
      float32 float64
    
      complex64 complex128
       ~~~
  > int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
  
  > int 和 float在进行运算时结果会有所不同，应当格外注意
  - 变量默认值
    > 数值类型为 0，
    布尔类型为 false，
    字符串为 "" (空字符串)
  - 类型转换(数值类型，需要显示声明(强制))
     ~~~go 
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)
     ~~~
  - 常量用const声明
   > 大写字母开头的常量在包外可见