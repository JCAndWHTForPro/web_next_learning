## math函数里面针对最大、最小、正无穷、负无穷、非法字符的表示
好的，你问到了 Go 语言中处理数学边界和特殊值的核心知识点。这在进行科学计算、数据分析或任何需要严谨浮点数处理的场景中都至关重要。

Go 遵循的是 **IEEE 754** 浮点数标准，这个标准定义了如何表示正无穷、负无穷和 "NaN" (Not a Number)。Go 的 `math` 包为我们提供了访问和检查这些特殊值的便捷方式。

---

### 1. 如何表示和获取这些特殊值？

Go `math` 包中预定义了这些特殊值的常量。它们都是 `float64` 类型。

#### a. 正无穷 (Positive Infinity)

*   **常量**: `math.Inf(1)`
*   **说明**: `math.Inf()` 函数接收一个 `int` 类型的符号参数。
    *   传入 `1` 或任何正数，返回正无穷 `+Inf`。
    *   传入 `-1` 或任何负数，返回负无穷 `-Inf`。
*   **生成方式**:
    *   `1.0 / 0.0` 会产生正无穷。
    *   一个非常大的数溢出时（如 `math.MaxFloat64 * 2`）。

**示例代码**:

```go
import "fmt"
import "math"

posInf := math.Inf(1)
fmt.Println(posInf) // 输出: +Inf

// 通过计算得到
z := 0.0
fmt.Println(1.0 / z) // 输出: +Inf
```

#### b. 负无穷 (Negative Infinity)

*   **常量**: `math.Inf(-1)`
*   **说明**: 与正无穷类似，只是符号参数为负。
*   **生成方式**:
    *   `-1.0 / 0.0` 会产生负无穷。
    *   一个非常小的负数下溢时。

**示例代码**:

```go
negInf := math.Inf(-1)
fmt.Println(negInf) // 输出: -Inf

z := 0.0
fmt.Println(-1.0 / z) // 输出: -Inf
```

#### c. 非法字符/非数字 (NaN - Not a Number)

*   **常量**: `math.NaN()`
*   **说明**: NaN 用于表示一个未定义或无法表示的数值结果，比如 "0 除以 0"。
*   **生成方式**:
    *   `0.0 / 0.0` 会产生 NaN。
    *   对负数开平方根 `math.Sqrt(-1.0)`。
    *   任何涉及 NaN 的算术运算，结果都是 NaN。

**示例代码**:

```go
notANumber := math.NaN()
fmt.Println(notANumber) // 输出: NaN

z := 0.0
fmt.Println(z / z) // 输出: NaN

fmt.Println(math.Sqrt(-1)) // 输出: NaN
```

#### d. 最大值 (Max Value) 和 最小值 (Min Value)

这些不是无穷大，而是 `float64` 或 `float32` 能表示的**有限**的最大/最小值。

*   **`float64` 的最大值**: `math.MaxFloat64`
*   **`float64` 的最小正数**: `math.SmallestNonzeroFloat64` (最接近0的正数)

*   **`float32` 的最大值**: `math.MaxFloat32`
*   **`float32` 的最小正数**: `math.SmallestNonzeroFloat32`

**示例代码**:

```go
fmt.Println(math.MaxFloat64) // 输出: 1.7976931348623157e+308
```

---

### 2. 如何检查一个值是否是这些特殊值？

这是非常重要的一步，因为你不能直接用 `==` 来判断。特别是 NaN，它有一个奇特的性质。

#### `NaN` 的奇特性质

**任何值与 `NaN` 进行 `==` 比较，结果永远是 `false`，即使是 `NaN == NaN` 也是 `false`！**

这是 IEEE 754 标准的规定。因为 NaN 代表的是“一个无法表示的结果”，而两个无法表示的结果并不一定相等（比如 `0/0` 的结果和 `sqrt(-1)` 的结果都是 NaN，但它们来源不同）。

所以，**绝对不能**用 `x == math.NaN()` 来检查！

#### 正确的检查方法

`math` 包提供了一系列 `Is...` 函数来进行可靠的检查。

*   **检查是否为 NaN**: `math.IsNaN(f float64) bool`
    *   这是**唯一正确**的检查 NaN 的方法。

*   **检查是否为无穷**: `math.IsInf(f float64, sign int) bool`
    *   `sign > 0`：检查是否为**正无穷**。
    *   `sign < 0`：检查是否为**负无穷**。
    *   `sign == 0`：检查是正无穷**或**负无穷。

**示例代码**:

```go
v1 := math.NaN()
v2 := math.Inf(1)
v3 := math.Inf(-1)
v4 := 123.45

// 检查 NaN
fmt.Printf("Is v1 NaN? %t\n", math.IsNaN(v1)) // 输出: Is v1 NaN? true
fmt.Printf("Is v4 NaN? %t\n", math.IsNaN(v4)) // 输出: Is v4 NaN? false

// 错误示范
if v1 == math.NaN() {
    fmt.Println("This will never be printed!")
}

// 检查无穷
fmt.Printf("Is v2 positive infinity? %t\n", math.IsInf(v2, 1))   // 输出: Is v2 positive infinity? true
fmt.Printf("Is v3 negative infinity? %t\n", math.IsInf(v3, -1))  // 输出: Is v3 negative infinity? true
fmt.Printf("Is v2 any infinity? %t\n", math.IsInf(v2, 0))     // 输出: Is v2 any infinity? true
fmt.Printf("Is v4 any infinity? %t\n", math.IsInf(v4, 0))     // 输出: Is v4 any infinity? false
```

---

### 总结表格

| 概念 | 如何获取/表示 | 如何检查 | 备注 |
| :--- | :--- | :--- | :--- |
| **正无穷 (+Inf)** | `math.Inf(1)` | `math.IsInf(val, 1)` | `1.0 / 0.0` 会产生 |
| **负无穷 (-Inf)** | `math.Inf(-1)` | `math.IsInf(val, -1)` | `-1.0 / 0.0` 会产生 |
| **任意无穷 (Inf)** | - | `math.IsInf(val, 0)` | 检查是正无穷或负无穷 |
| **非数字 (NaN)** | `math.NaN()` | `math.IsNaN(val)` | **唯一**的正确检查方法。`val == val` 自身比较为 `false` 是其特征。 |
| **最大 `float64`** | `math.MaxFloat64` | `val == math.MaxFloat64` | 有限的最大值 |
| **最小正数 `float64`**| `math.SmallestNonzeroFloat64` | `val == math.SmallestNonzeroFloat64`| 有限的、最接近0的正数 |

掌握这些函数的使用，可以让你编写出非常健壮、能够正确处理浮点数运算中各种边界情况的 Go 程序。


## 无类型常量的深入
好的，无类型常量（Untyped Constants）是 Go 语言一个非常独特且强大的特性，初看时确实容易让人困惑。它解决了其他静态语言（如 C++/Java）中一个常见的“烦恼”。

我们还是用一个简单的比喻来开始，然后深入细节。

---

### 比喻：一个“万能”的数字

想象一下，在 Go 语言里，当你写下一个**常量面值**（比如 `100` 或 `3.14`）时，这个常量在诞生之初，它并不是一个严格的 `int` 或 `float64`。

**你可以把它想象成一个“纯粹的、高精度的、万能的数学概念”**。

*   `const Pi = 3.14`：这里的 `Pi` 不是 `float64`，它就是数学概念上的“三点一四”，精度要多高有多高。
*   `const MaxSize = 1024`：这里的 `MaxSize` 不是 `int`，它就是数学概念上的“一千零二十四”。

这个“万能概念”非常灵活，它会**延迟（delay）**自己的类型确定，直到它被用在一个需要明确类型的上下文里。那时，它才会“变形”成最适合那个上下文的类型。

---

### 为什么需要这个特性？—— 解决“烦恼”

在 C++ 或 Java 里，你可能会遇到这样的问题：
```c++
const double PI = 3.14159;
int radius = 5;
double circumference = 2 * PI * radius; // 错误！不能用一个 double 和一个 int 直接相乘
```
你必须做显式的类型转换 `2 * PI * (double)radius`。这很繁琐。

Go 的无类型常量优雅地解决了这个问题。

### Go 的无类型常量如何工作？

我们来看 Go 的代码，并分析那个“万能概念”是如何“变形”的。

```go
package main

import "fmt"

// MaxSize 是一个无类型的整数常量。它只是数字 1024。
const MaxSize = 1024

// Pi 是一个无类型的浮点数常量。它只是数字 3.14。
const Pi = 3.14

func main() {
    // 场景一：MaxSize 遇到了一个需要 int 的地方
    var anInt int = MaxSize 
    // MaxSize 发现：“哦，右边需要一个 int，我 1024 可以完美地变成 int。”
    // 于是 MaxSize 在这里“变形”成了 int 类型。
    fmt.Printf("anInt: %T, %d\n", anInt, anInt) // 输出: anInt: int, 1024

    // 场景二：MaxSize 遇到了一个需要 float64 的地方
    var aFloat64 float64 = MaxSize
    // MaxSize 发现：“哦，右边需要一个 float64，我 1024 也可以完美地变成 float64。”
    // 于是 MaxSize 在这里“变形”成了 float64 类型。
    fmt.Printf("aFloat64: %T, %f\n", aFloat64, aFloat64) // 输出: aFloat64: float64, 1024.000000

    // 场景三：无类型常量之间的混合运算 (最神奇的地方!)
    // result := 2 * Pi * MaxSize // 假设有这么一个运算
    // 这里的 2、Pi、MaxSize 全都是无类型的“万能概念”。
    // Go 会用非常高的精度来计算它们的结果。
    // 只有在最后需要把结果赋值给一个有类型的变量时，
    // 最终结果才会“变形”并接受类型检查。

    var someResult float64 = 2 * Pi * MaxSize
    // Go 内部计算: 2 * 3.14 * 1024 = 6530.56 (高精度计算)
    // 然后发现需要赋值给一个 float64，于是 6530.56 变形为 float64。
    fmt.Printf("someResult: %T, %f\n", someResult, someResult) // 输出: someResult: float64, 6530.560000

    // 场景四：如果“变形”失败会怎样？
    // var anInt8 int8 = MaxSize // MaxSize 是 1024
    // MaxSize 发现：“哦，右边需要一个 int8，但 int8 的最大值是 127。”
    // “我 1024 太大了，塞不进去！变形失败！”
    // 这行代码会导致编译错误: constant 1024 overflows int8
}
```

### 有类型常量 (Typed Constants)

你也可以在声明常量时就给它一个明确的类型，剥夺它的“万能性”。

```go
const TypedPi float64 = 3.14 // TypedPi 从诞生起就是一个 float64
const TypedSize int = 1024     // TypedSize 从诞生起就是一个 int
```
现在，它们就和普通的变量一样，遵循严格的类型匹配规则。

```go
var anotherFloat float64 = TypedSize // 错误！不能把一个 int 类型的常量赋值给 float64 变量
                                    // 必须写成 float64(TypedSize)
```

### 总结：无类型常量的核心优势

1.  **灵活性 (Flexibility)**：它们可以自由地与不同类型的变量混合使用，只要它们的值在那个类型的表示范围内。这大大减少了不必要的显式类型转换，让代码更简洁。
2.  **高精度 (High Precision)**：在常量表达式的计算过程中，Go 会使用比标准类型（如 `float64`）更高的精度来计算，只有在最后赋值时才会发生可能的精度损失。这避免了中间计算步骤的误差累积。

### 概括的表格

| 概念 | 简单概括 | Go 代码示例 |
| :--- | :--- | :--- |
| **无类型常量** (Untyped Constant) | 一个**高精度的、临时的“万能数学概念”**，它会**延迟**确定自己的类型，直到被用在一个需要明确类型的上下文中。 | `const N = 10` (N 可以是 int, int8, float64...) |
| **有类型常量** (Typed Constant) | 从声明的那一刻起，就**被赋予了明确的类型**，失去了“万能性”，遵循严格的类型规则。 | `const N int = 10` (N 永远是 int) |
| **默认类型** | 如果一个无类型常量最终没有被明确的上下文赋予类型（比如直接打印它），它会有一个**默认类型**：`bool`, `rune`, `int`, `float64`, `complex128`, 或 `string`。| `fmt.Printf("%T", 10)` -> `int`<br>`fmt.Printf("%T", 3.14)` -> `float64` |

**简单记住**：Go 的无类型常量就像一团“橡皮泥”，在放入一个模子（有类型的变量）之前，它可以是任何形状。而有类型常量就像一块“砖头”，形状从一开始就固定了。