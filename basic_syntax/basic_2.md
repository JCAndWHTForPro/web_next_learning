## 2025-08-03

1. 接口的相关能力

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


2. interface的特殊用法：代表任意类型
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

3. interface特殊用法：类型断言

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

4. interface特殊用法：switch类型判断
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

5. 我两个interface，一样的方法，实现的struct的实例问题
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

