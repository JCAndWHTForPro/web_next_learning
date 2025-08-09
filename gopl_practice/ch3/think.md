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