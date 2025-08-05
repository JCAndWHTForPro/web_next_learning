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

5. 常量注意的点
```go
1、常量中中使用的函数必须是内置函数，否则编译不过
package main

import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c)
}

2、c = unsafe.Sizeof(a)的结果是16 的原因：

Go 中的字符串类型本质上是一个结构体，由两个字段组成：
一个指向底层字节数组的指针（*byte）
一个表示字符串长度的整数（int）
在 64 位系统中：
指针类型（*byte）占用 8 字节
整数类型（int）占用 8 字节
两者相加正好是 16 字节，因此 unsafe.Sizeof(a) 返回 16。

3、常量可以不用赋值，下面的赋值操作都是重复上面的
package main

import "fmt"
const (
i=1<<iota
j=3<<iota
k
l
)

func main() {
fmt.Println("i=",i)
fmt.Println("j=",j)
fmt.Println("k=",k)
fmt.Println("l=",l)
}
```

6. 注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。


7. swith的骚操作
```go
package main

import "fmt"
/**
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

Type Switch 语法格式如下：

switch x.(type){
    case type:
       statement(s);
    case type:
       statement(s);
    /* 你可以定义任意个数的case */
default: /* 可选 */
statement(s);
}
 */
func main() {

    switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }
}
```

8. swith默认是有break的，这个和其他语言不一样，要想没有要加东西
```go
/**
使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
 */
package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
```

9. 通道语句select
```go
elect {
  case <- channel1:
    // 执行的代码
  case value := <- channel2:
    // 执行的代码
  case channel3 <- value:
    // 执行的代码

    // 你可以定义任意数量的 case

  default:
    // 所有通道都没有准备好，执行的代码
}
```
以下描述了 select 语句的语法：

- 每个 case 都必须是一个通道
- 所有 channel 表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通道可以进行，它就执行，其他被忽略。
- 如果有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行。
  否则：
  - 如果有 default 子句，则执行该语句。
  - 如果没有 default 子句，select 将阻塞，直到某个通道可以运行；Go 不会重新对 channel 或值进行求值。
  实例


10. 函数与方法
```
text简言之，方法是 “绑定了类型的函数”，函数强调通用性，方法强调类型的专属行为。
```
```go
package main

import "fmt"

// 定义一个结构体（自定义类型）
type Person struct {
    name string
    age  int
}

// 函数：独立功能，格式化字符串
func formatInfo(name string, age int) string {
    return fmt.Sprintf("姓名：%s，年龄：%d", name, age)
}

// 方法：为 Person 类型定义行为
func (p Person) introduce() string {
    // 方法内可直接访问 Person 的私有字段 name 和 age
    return formatInfo(p.name, p.age) // 调用函数
}

func main() {
    p := Person{name: "Alice", age: 30}
    
    // 调用方法（通过接收者 p）
    fmt.Println(p.introduce()) // 输出：姓名：Alice，年龄：30
    
    // 调用函数（直接调用）
    fmt.Println(formatInfo("Bob", 25)) // 输出：姓名：Bob，年龄：25
}
```

11. 全局变量的使用

```go
// 包名（非main包）
package mypackage

// 1. 首字母大写的全局变量（可导出，其他包可访问）
var PublicVar int = 100
var AppName string = "MyApp"

// 2. 首字母小写的全局变量（不可导出，仅当前包内可见）
var privateVar string = "仅内部使用"
    
```
```go
package main

import (
    "fmt"
    "your-project-path/mypackage" // 导入自定义包
)

func main() {
    // 访问导出的全局变量（包名.变量名）
    fmt.Println(mypackage.PublicVar) // 输出：100
    fmt.Println(mypackage.AppName)   // 输出：MyApp

    // 以下代码会编译错误（私有变量不可访问）
    // fmt.Println(mypackage.privateVar) 
}
    
```
```text
注意：要使用
go mod init
初始化你的项目，才能进行包名的导入用，否则导入不了。这个和java还不太一样
或者是直接在文件夹下面创建一个go的模块文件
```


12. 数组与切片的区别

```go
// 这种是切片
var arr = []int{123}
// 这种是数组
var arr = [1]int{12}
```

| 特性               | 切片 `[]int{123}`                  | 数组 `[1]int{12}`                  |
|--------------------|------------------------------------|------------------------------------|
| 长度是否固定       | 动态可变（可通过 `append` 扩容）   | 固定不变（长度是类型的一部分）     |
| 类型表示           | `[]int`（不包含长度）              | `[1]int`（包含长度 `1`）           |
| 内存分配           | 切片本身是引用类型，数据存在底层数组（可能在栈或堆上） | 数组是值类型，数据直接存储在数组本身 |
| 赋值/传参行为      | 传递的是切片引用（不复制底层数据） | 传递的是数组副本（复制全部元素）   |
| 扩容机制           | 超过容量时会自动扩容（创建新底层数组） | 无扩容机制，长度固定               |

```go
package main

import "fmt"

func main() {
    // 切片
    s := []int{123}
    fmt.Printf("切片类型: %T, 长度: %d, 容量: %d\n", s, len(s), cap(s)) // 输出：[]int, 1, 1
    
    // 数组
    a := [1]int{12}
    fmt.Printf("数组类型: %T, 长度: %d\n", a, len(a)) // 输出：[1]int, 1

    // 切片可以动态扩容
    s = append(s, 456)
    fmt.Printf("扩容后切片: 长度=%d, 容量=%d\n", len(s), cap(s)) // 输出：2, 2（容量可能自动翻倍）

    // 数组长度固定，无法直接添加元素
    // a = append(a, 34) // 编译错误：first argument to append must be slice
}
```

13. c语言体系中经典问题[3]*int和*[3]int

| 类型          | 含义解释                                                                 | 特性                                                                 |
|---------------|--------------------------------------------------------------------------|----------------------------------------------------------------------|
| `[3]*int`     | 一个**数组**（长度为 3），数组中的每个元素都是 `*int` 类型的指针          | 1. 本质是数组（值类型）<br>2. 数组本身存储的是指针（共 3 个指针）<br>3. 赋值/传参会复制整个数组（复制 3 个指针） |
| `*[3]int`     | 一个**指针**，指向一个长度为 3 的 `int` 类型数组                          | 1. 本质是指针（引用类型）<br>2. 指针指向一个包含 3 个 `int` 的数组<br>3. 赋值/传参仅复制指针（不复制数组数据） |

```go
package main

import "fmt"

func main() {
    // [3]*int：数组元素是指针
    var a [3]*int
    num := 10
    a[0] = &num  // 给数组第一个元素赋值（存指针）
    fmt.Println(a[0])  // 输出指针地址：0xc00001a0a8

    // *[3]int：指向数组的指针
    var b [3]int = [3]int{1, 2, 3}
    var p *[3]int = &b  // 指针指向数组 b
    fmt.Println(p[0])   // 输出数组第一个元素：1（通过指针访问数组）
}
```
```go
package main

import "fmt"

func main() {
	// 这种不是数组，是一个切片
	var arr = []int{}
	arr = append(arr, 123)
	fmt.Println(arr)

	// 这种是数组
	var arr_real = [3]int{123, 234, 345}
	fmt.Println(arr_real)

	// 这种是数组指针
	var arr_ptr *[3]int
	arr_ptr = &arr_real
	fmt.Println(*arr_ptr)

	// 这两种使用指针访问数组的方式是等价的，上面是语法糖
	arr_ptr[1] = 999
	(*arr_ptr)[1] = 101010
	fmt.Println(arr_ptr)

}

```


14. go里面针对结构体（其实就是c里面的）的一些思索
```go

// 使用指针类型，修改会影响原变量
func modifyPointer(sw *StringWriter) {
    /**
        两种写法效果完全相同，这种是go的语法糖
	当你有一个结构体指针 sw *StringWriter 时，
	Go 允许你直接使用 sw.buffer 来访问结构体字段，
	而无需显式写成 (*sw).buffer。
	编译器会自动帮你完成指针解引用的操作，
	这是 Go 为了简化代码而设计的语法规则。两种写法在语义上完全等价
     */
    sw.buffer = "modified"
    // (*sw).buffer = "modified"
}

```

