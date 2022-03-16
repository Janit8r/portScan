package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func insertDB(ip string, port string) {
	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	defer db.Close()
	sqlStr := "insert into ipInfo (ip,port) values(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("预处理出问题:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(ip, port)
	if err != nil {
		fmt.Printf("插入数据失败:%v", err)
	}
	fmt.Println("数据插入成功")
	return
}
func selectDB(ip string) string {
	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	defer db.Close()
	sqlStr := "select * from ipInfo where ip=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("预处理出问题:%v\n", err)
		return "err"
	}
	defer stmt.Close()
	rows, err := stmt.Query(ip)
	if err != nil {
		fmt.Printf("查询失败")
		return "query error"
	}
	defer rows.Close()

	var port string
	for rows.Next() {
		rows.Scan(&ip, &port)

	}

	return ip + "-" + port
}
