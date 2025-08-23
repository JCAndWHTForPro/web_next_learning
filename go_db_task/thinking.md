## 一、Sqlx入门



Sqlx 是 Go 语言中一个非常流行的数据库操作库，它在标准库 `database/sql` 的基础上提供了更便捷的 API，支持结构体与数据库记录的映射，减少了重复代码。下面是 Sqlx 的入门指南，包括安装、基本使用和核心功能。


### 1. 安装 Sqlx
首先需要安装 Sqlx 库，使用 `go get` 命令：
```bash
go get github.com/jmoiron/sqlx
```

同时，你还需要安装对应数据库的驱动（以 MySQL 为例）：
```bash
go get github.com/go-sql-driver/mysql
```


### 2. 基本使用流程
#### 步骤 1：连接数据库
使用 `sqlx.Connect()` 建立数据库连接，参数与 `database/sql` 相同：

```go
package main

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql" // MySQL 驱动
)

// 定义一个结构体映射数据库表
type User struct {
    ID   int    `db:"id"`    // db 标签指定数据库字段名
    Name string `db:"name"`
    Age  int    `db:"age"`
}

func main() {
    // 连接数据库（DSN 格式：用户名:密码@tcp(地址:端口)/数据库名?参数）
    db, err := sqlx.Connect("mysql", "root:password@tcp(127.0.0.1:3306)/testdb?parseTime=true")
    if err != nil {
        fmt.Printf("数据库连接失败: %v\n", err)
        return
    }
    defer db.Close() // 程序结束时关闭连接

    // 测试连接是否有效
    if err := db.Ping(); err != nil {
        fmt.Printf("连接测试失败: %v\n", err)
        return
    }
    fmt.Println("数据库连接成功")
}
```


#### 步骤 2：执行 SQL 操作
Sqlx 提供了与 `database/sql` 类似的方法（如 `Exec`、`Query` 等），但增加了对结构体的支持。

##### （1）插入数据（INSERT）
```go
// 插入单条数据
func insertUser(db *sqlx.DB) {
    sql := "INSERT INTO users (name, age) VALUES (?, ?)"
    result, err := db.Exec(sql, "Alice", 25)
    if err != nil {
        fmt.Printf("插入失败: %v\n", err)
        return
    }
    // 获取插入的 ID
    id, _ := result.LastInsertId()
    fmt.Printf("插入成功，ID: %d\n", id)
}
```

##### （2）查询单条数据（SELECT）
使用 `Get()` 方法将查询结果映射到结构体：
```go
// 查询单条数据
func getUser(db *sqlx.DB, id int) {
    var user User
    // 注意：SQL 中的字段名需与结构体的 db 标签对应
    sql := "SELECT id, name, age FROM users WHERE id = ?"
    err := db.Get(&user, sql, id)
    if err != nil {
        fmt.Printf("查询失败: %v\n", err)
        return
    }
    fmt.Printf("查询结果: %+v\n", user) // %+v 显示字段名
}
```

##### （3）查询多条数据（SELECT）
使用 `Select()` 方法将结果映射到结构体切片：
```go
// 查询多条数据
func getUsers(db *sqlx.DB) {
    var users []User
    sql := "SELECT id, name, age FROM users WHERE age > ?"
    err := db.Select(&users, sql, 18)
    if err != nil {
        fmt.Printf("查询失败: %v\n", err)
        return
    }
    fmt.Printf("查询结果: %+v\n", users)
}
```

##### （4）更新数据（UPDATE）
```go
// 更新数据
func updateUser(db *sqlx.DB, id int, newAge int) {
    sql := "UPDATE users SET age = ? WHERE id = ?"
    result, err := db.Exec(sql, newAge, id)
    if err != nil {
        fmt.Printf("更新失败: %v\n", err)
        return
    }
    // 获取受影响的行数
    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("更新成功，影响行数: %d\n", rowsAffected)
}
```

##### （5）删除数据（DELETE）
```go
// 删除数据
func deleteUser(db *sqlx.DB, id int) {
    sql := "DELETE FROM users WHERE id = ?"
    result, err := db.Exec(sql, id)
    if err != nil {
        fmt.Printf("删除失败: %v\n", err)
        return
    }
    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("删除成功，影响行数: %d\n", rowsAffected)
}
```


### 3. 核心功能与优势
#### （1）结构体映射
通过 `db:"字段名"` 标签，Sqlx 能自动将查询结果映射到结构体，避免手动赋值：
```go
type User struct {
    ID   int    `db:"id"`   // 对应数据库的 id 字段
    Name string `db:"name"` // 对应数据库的 name 字段
    Age  int    `db:"age"`  // 对应数据库的 age 字段
}
```

#### （2）命名参数查询
支持使用 `:参数名` 形式的命名参数，提高 SQL 可读性：
```go
sql := "SELECT * FROM users WHERE name = :name AND age > :age"
params := map[string]interface{}{
    "name": "Bob",
    "age":  20,
}
var users []User
err := db.NamedQuery(&users, sql, params) // 注意：NamedQuery 需配合结构体指针使用
```

#### （3）事务支持
Sqlx 对事务的支持与标准库类似，但提供了 `MustBegin()` 等便捷方法：
```go
// 开启事务
tx := db.MustBegin() // 简化版，出错会 panic

// 执行事务操作
_, err := tx.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Charlie", 30)
if err != nil {
    tx.Rollback() // 出错回滚
    return
}

// 提交事务
err = tx.Commit()
if err != nil {
    fmt.Printf("事务提交失败: %v\n", err)
}
```

#### （4）批量操作
使用 `NamedExec` 批量插入数据：
```go
users := []User{
    {Name: "Dave", Age: 28},
    {Name: "Eve", Age: 22},
}
sql := "INSERT INTO users (name, age) VALUES (:name, :age)"
_, err := db.NamedExec(sql, users) // 批量插入
```


### 4. 注意事项
1. **数据库驱动**：Sqlx 依赖标准库的 `database/sql`，需导入对应数据库的驱动（如 MySQL、PostgreSQL 等），但不需要直接使用驱动的 API。
2. **结构体标签**：`db` 标签必须与数据库字段名一致，否则映射会失败（可通过 `db:"-"` 忽略不需要的字段）。
3. **错误处理**：Sqlx 提供了 `Must` 前缀的方法（如 `MustConnect`、`MustBegin`），出错时会直接 `panic`，适合快速开发；生产环境建议显式处理错误。
4. **性能**：Sqlx 的性能与 `database/sql` 接近，额外开销主要来自结构体反射映射，适合大多数场景。


通过以上内容，你可以快速入门 Sqlx 的基本使用。更多高级功能（如自定义映射、JSON 支持等）可参考 [Sqlx 官方文档](https://github.com/jmoiron/sqlx)。