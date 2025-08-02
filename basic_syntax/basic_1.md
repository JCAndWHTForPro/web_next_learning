## 2025-08-02
1. go的基础问题：错误: 无法在构建后运行,'main' 文件具有非 main 软件包或不包含 'main' 函数
```text
这个错误通常出现在 Go 语言项目中，原因是程序入口不符合 Go 的执行规范。Go 语言要求：
程序的入口文件必须声明为 main 包（package main）
入口文件中必须包含 main 函数（func main()）
```

2. 声明中的语法注意
```text
intVal := 1 相等于：

var intVal int 
intVal =1 
```
我们知道可以在变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 var 关键字其实是显得有些多余了，因此我们可以将它们简写为 a := 50 或 b := false。

```text
a 和 b 的类型（int 和 bool）将由编译器自动推断。
这是使用变量的首选形式，但是它只能被用在函数体内，
而不可以用于全局变量的声明与赋值。
使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
```

3. 推断赋值语句的使用：

```text
右边的这些值以相同的顺序赋值给左边的变量，所以 a 的值是 5， b 的值是 7，c 的值是 "abc"。
这被称为 并行 或 同时 赋值。
如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。
```

4. 全局变量的声明
```go
// 自定义包（非main包）
package mypackage

// 1. 可导出的全局变量（首字母大写，其他包可访问）
var PublicVar string = "我是跨包可见的全局变量"
var Version int = 1.0

// 2. 不可导出的全局变量（首字母小写，仅当前包内可见）
var privateVar int = 100
var config map[string]string = map[string]string{
    "mode": "dev",
}

// 3. 批量声明全局变量（混合可导出和不可导出）
var (
    MaxSize  int = 1024  // 可导出
    minSize  int = 16    // 不可导出
    isDebug  bool = true // 不可导出
)

// 包内函数可以访问所有全局变量（无论是否导出）
func PrintVars() {
    println("私有变量:", privateVar)
    println("配置:", config["mode"])
    println("最小尺寸:", minSize)
}

```