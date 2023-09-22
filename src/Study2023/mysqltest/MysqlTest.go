package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "gorm.io/gorm/logger"
)

type CustomerBind struct {
	Id          string
	MainEcifId  string
	SubEcifId   string
	CreateTime  string
	OperateTime sql.NullString
}

/*
*

I.异常1:MysqlTest.go:7:2: no required module provides package github.com/go-sql-driver/mysql; to add it:
解决:
你在使用 github.com/go-sql-driver/mysql 包之前需要先安装它。你可以通过以下命令使用 Go Modules 安装该依赖项：

go get -u github.com/go-sql-driver/mysql

I.异常2 ：log
go get gorm.io/gorm/logger
*/
func main() {
	fmt.Println("开始")
	db1()
	db2()
}

func db1() {
	// 连接数据库
	db, err := sql.Open("mysql", "fso_cromwell:fso_cromwell8K0STD6ddev@tcp(dev17714.db.test:4301)/fso_cromwell")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 查询数据
	rows, err := db.Query("SELECT * FROM t_customer_bind")
	if err != nil {
		fmt.Println("查询数据失败:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var main_ecif_id string
		var sub_ecif_id string
		var sub_ecif_id1 string
		var sub_ecif_id2 sql.NullString

		err = rows.Scan(&id, &main_ecif_id, &sub_ecif_id, &sub_ecif_id1, &sub_ecif_id2)
		if err != nil {
			fmt.Println("获取数据失败:", err)
			return
		}
		fmt.Printf("ID: %s, Name: %s, Age: %s Age3: %s Age4: %s \n", id, main_ecif_id, sub_ecif_id, sub_ecif_id1, sub_ecif_id2)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("遍历结果集失败:", err)
		return
	}
}

func db2() {

	db, err := sql.Open("mysql", "fso_cromwell:密码@tcp(xxxxx:4301)/fso_cromwell")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM t_customer_bind")
	if err != nil {
		fmt.Println("查询数据失败:", err)
		return
	}
	defer rows.Close()

	var customerBindList []CustomerBind

	for rows.Next() {
		var customerBind CustomerBind

		err = rows.Scan(&customerBind.Id, &customerBind.MainEcifId, &customerBind.SubEcifId, &customerBind.CreateTime, &customerBind.OperateTime)
		if err != nil {
			fmt.Println("获取数据失败:", err)
			return
		}
		fmt.Printf("ID: %s   \n", customerBind)
		customerBindList = append(customerBindList, customerBind)
	}

	for t := range customerBindList {
		fmt.Println("打印结果集:", customerBindList[t])
	}

	for t := range customerBindList {
		if customerBindList[t].OperateTime.Valid {
			fmt.Println("打印结果集  有值:", customerBindList[t].OperateTime.String)
		} else {
			fmt.Println("打印结果集  无值:", customerBindList[t].OperateTime.String)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("遍历结果集失败:", err)
	}

}
