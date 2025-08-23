package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Grade int    `db:"grade"`
}

func main() {
	connect, err := sqlx.Connect("mysql", "root:JI-109385147-ch@tcp(47.92.123.15:3306)/go_lesson_db")
	if err != nil {
		fmt.Printf("数据库连接失败：%v\n", err)
		return
	}
	fmt.Println("数据库连接成功")
	defer connect.Close()
	insertOneInfo(connect)
	updateInfoByName(connect, "张三", 4)
	selectByAge(connect, ">", 18)

}
