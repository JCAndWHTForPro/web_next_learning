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