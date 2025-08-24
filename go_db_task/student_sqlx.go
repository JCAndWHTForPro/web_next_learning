package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

/*
*
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。

- 要求 ：
  - 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
  - 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
  - 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
  - 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

（Sqlx入门的题目1和题目2，proctice的点都一样，就不做练习）
*/
func insertOneInfo(conn *sqlx.DB) int64 {
	sql := "insert into students (name,age,grade) values(?,?,?)"
	exec, err := conn.Exec(sql, "张三", 20, 3)
	if err != nil {
		fmt.Printf("执行sql失败：%v\n", err)
		panic("失败了啊")
	}
	id, _ := exec.LastInsertId()
	fmt.Println("执行数据库成功，id是：", id)
	return id
}
func updateInfoByName(conn *sqlx.DB, name string, grade int) bool {
	sql := "update students set grade = ? where name = ?"
	rst, err := conn.Exec(sql, grade, name)
	if err != nil {
		fmt.Printf("执行update sql失败：%v\n", err)
		return false
	}
	affected, _ := rst.RowsAffected()
	fmt.Println("更新成功，更新的行数是：", affected)
	return true
}

func selectByAge(coon *sqlx.DB, operation string, age int) []User {
	var users []User
	var sql string
	if operation == ">" {
		sql = "select id,name,age,grade from students where age > ?"
	} else if operation == "<" {
		sql = "select id,name,age,grade from students where age < ?"
	} else {
		fmt.Println("暂时不支持这个操作")
		return users
	}
	err := coon.Select(&users, sql, age)
	if err != nil {
		fmt.Printf("查询失败：%v\n", err)
		return users
	}
	fmt.Println("查询成功", users)
	return users
}
