### 第三章节
- 指针
  - \指针的定义 var varName *varType
     > varType可以是指针
     ~~~go
     // 定义了一个名为intPointer类型为int的指针
     var intPointer *int 
     // 定义了一个名为strPointer类型为string的指针
     var strPointer = *string
    
     // 定义了一个名为arrayPointer类型为string的指针数组
     var arrayPointer = []*string
     ...
     ~~~
  - 指针其实就是指向某个变量的内存地址的变量
    > 取地址 &var  
      ~~~go
      str := "hello go"
      将指针strPointer指向变量str的地址：
      strPointer := &str
      ~~~
    
    > 取值 *varName
      ~~~go
      // fmt.Println(*strPointer)
      // 最终会打印 ”hello go“
      ~~~
  - 指针和变量共享值，其作用就目前阶段来说就是使得在值传递的过程中减少内存消耗（没有值的拷贝）
     > 如果改变了某个变量的值，那么只要指向了这个变量或者是间接指向了这个变量的指针也会随之变量值的改变而改变
  - 指针本质上是一个变量，不能指向常量的地址！！！