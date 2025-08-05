## 2025-08-03

### 1. 接口的相关能力

- Go 中没有关键字显式声明某个类型实现了某个接口。
- 只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口。
- 接口变量可以存储实现该接口的任意值。
- 接口变量实际上包含了两个部分：
  - 动态类型：存储实际的值类型。
  - 动态值：存储具体的值。
- 接口的零值是 nil。
- 一个未初始化的接口变量其值为 nil，且不包含任何动态类型或值。
- 定义为 interface{}，可以表示任何类型。

```text
1、多态：不同类型实现同一接口，实现多态行为。
2、解耦：通过接口定义依赖关系，降低模块之间的耦合。
3、泛化：使用空接口 interface{} 表示任意类型。
```
```go
/* 定义接口 */
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

/* 定义结构体 */
type struct_name struct {
   /* variables */
}

/* 实现接口方法 */
func (struct_name_variable struct_name) method_name1() [return_type] {
   /* 方法实现 */
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
   /* 方法实现*/
}
```


### 2. interface的特殊用法：代表任意类型
```go

import "fmt"

func printValueType(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

/*
*
空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。
1、任意类型都实现了空接口。
2、常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。
*/
func main() {
	printValueType(11)
	printValueType("string")
	printValueType(12.234)
	printValueType([3]int{123, 234, 345})
	printValueType([]int{123, 234, 345})
}

```

### 3. interface特殊用法：类型断言

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{} = 123
	a, er := val.(string)
	if er {
		fmt.Println(a)
		fmt.Println(reflect.TypeOf(val))
	} else {
		fmt.Println("不是预想的类型")
		fmt.Println(reflect.TypeOf(val))
	}
}

```

### 4. interface特殊用法：switch类型判断
```go
package main

import "fmt"

func printType(val interface{}) {
        switch v := val.(type) {
        case int:
                fmt.Println("Integer:", v)
        case string:
                fmt.Println("String:", v)
        case float64:
                fmt.Println("Float:", v)
        default:
                fmt.Println("Unknown type")
        }
}

func main() {
        printType(42)
        printType("hello")
        printType(3.14)
        printType([]int{1, 2, 3})
}
```

### 5. 我两个interface，一样的方法，实现的struct的实例问题
```text
在 Go 中，接口的实现是隐式的，只要一个类型实现了接口的所有方法，
它就会自动满足该接口。如果两个接口的方法完全相同，
那么实现这些方法的结构体会同时满足这两个接口。
```
```go
package main

import "fmt"

// 定义两个接口，方法完全相同
type A interface {
    Method()
}

type B interface {
    Method()
}

// 实现 Method 的结构体
type S struct{}

func (s S) Method() {
    fmt.Println("S.Method()")
}

func main() {
    var a A = S{}  // 合法：S 实现了 A
    var b B = S{}  // 合法：S 实现了 B

    a.Method() // 输出: S.Method()
    b.Method() // 输出: S.Method()
}

```
### 6. defer语法
```text
在 Go 语言中，defer 关键字用于延迟执行函数调用，
它会在当前函数（或代码块）执行结束前（返回前）被触发，
无论函数是正常返回还是因 panic 退出。
这一特性非常适合处理资源释放、清理操作等场景。
```
```go
func main() {
    fmt.Println("start")
    
    // 延迟执行下面的打印语句
    defer fmt.Println("deferred")
    
    fmt.Println("end")
}
```
核心的特性：
1. 延迟执行时机：defer 后的函数会在包含它的函数即将返回时执行，无论是否发生错误或 panic。
2. 多个 defer 的执行顺序： 多个 defer 按后进先出（LIFO） 顺序执行（类似栈）：

```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}

/**
输出：
3
2
1
 */
```

3. 参数即时求值：defer 函数的参数在声明时就会被计算，而非执行时：
```go
func main() {
    i := 0
    defer fmt.Println(i)  // 参数 i 此时为 0
    i = 100
}
```
4. 典型应用场景
   1. 资源释放:用于关闭文件、网络连接、数据库连接等，确保资源被释放：
        ```go
        func readFile(path string) {
            file, err := os.Open(path)
            if err != nil {
                return
            }
            defer file.Close()  // 函数结束时自动关闭文件
            
            // 读取文件操作...
        }
        ```
   2. 锁定与解锁:配合互斥锁（sync.Mutex）使用，避免忘记解锁：
        ```go
        var mu sync.Mutex
        var count int
        
        func increment() {
            mu.Lock()
            defer mu.Unlock()  // 函数结束时自动解锁
            count++
        }
        ```
   3. 错误处理与日志记录:记录函数执行耗时或错误信息：
        ```go
        func doTask() {
            start := time.Now()
            defer func() {
                fmt.Printf("任务耗时: %v\n", time.Since(start))
            }()  // 匿名函数作为defer
            
            // 执行任务...
        }
        ```
5. 注意事项
- defer 会略微影响性能（尤其在循环中大量使用时），但可读性和安全性更重要。
- defer 可以用于匿名函数，适合需要复杂逻辑的延迟操作。
- 在 return 语句中，defer 会在返回值确定后、函数退出前执行（可能修改命名返回值）：
    ```go
    func f() (result int) {
        defer func() { result++ }()
        return 0  // 实际返回 1
    }
    ```


### 7. 数组的操作
在 Go 中，数组（array）可以进行类似切片（slice）的切割操作， 但结果会返回一个切片（slice） 而非新数组。数组切割使用 [start:end] 语法，规则与切片切割一致：
- start：起始索引（包含），默认为 0
- end：结束索引（不包含），默认为数组长度
- 结果是一个指向原数组的切片，共享底层数据
```go
package main

import "fmt"

func main() {
    // 定义一个长度为5的int数组
    arr := [5]int{10, 20, 30, 40, 50}
    
    // 切割数组，生成切片
    s1 := arr[1:3]   // 从索引1到3（不包含3），结果：[20, 30]
    s2 := arr[:2]    // 从开头到索引2，结果：[10, 20]
    s3 := arr[3:]    // 从索引3到结尾，结果：[40, 50]
    s4 := arr[:]     // 整个数组转为切片，结果：[10,20,30,40,50]
    
    fmt.Println(s1) // [20 30]
    fmt.Println(s2) // [10 20]
    fmt.Println(s3) // [40 50]
    fmt.Println(s4) // [10 20 30 40 50]
    
    // 注意：切割结果是切片（[]int），而非数组
    fmt.Printf("s1类型: %T\n", s1) // 输出：s1类型: []int
}
```
**关键注意点**
1. **结果类型**：数组切割后得到的是切片（slice），而非新数组。数组是固定长度的，而切片是动态长度的视图。
2. **数据共享**：切割生成的切片与原数组共享底层数据。修改切片会影响原数组，反之亦然：
```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:3]
s[0] = 200  // 修改切片元素
fmt.Println(arr) // 输出：[10 200 30 40 50]（原数组被修改）
```
总结：数组可以通过 [start:end] 语法切割，但结果是切片，且与原数组共享数据。这一点与切片切割的行为一致，但数组本身是固定长度的，无法直接改变长度。


### 8. 接口类型探索
**接口变量的工作方式是：**
- 它可以存储任何实现了该接口的具体类型的值（包括值类型或指针类型，如 StringWriter 或 *StringWriter）
- 接口变量本身就像一个 “容器”，内部存储了两部分信息：具体值的类型和具体值的数据
- &StringWriter{}（后面带大括号这种，其实就是初始化类结构体，然后取结构体对象的地址，就是指针）的本质，是一个StringWriter这个结构体的指针，就是*StringWriter

