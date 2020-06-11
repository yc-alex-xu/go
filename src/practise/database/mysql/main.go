package main

import (
	"database/sql"
	"fmt"
	"os"

	//"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:"+os.Getenv("MySQL_PASSWORD")+"@tcp("+os.Getenv("MySQL_HOST")+":3306)/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("john", "研发", "2012-12-09")
	res, err = stmt.Exec("smith", "财务", "2018-12-09")
	res, err = stmt.Exec("mary", "人事", "2020-12-09")
	res, err = stmt.Exec("kate", "销售", "2020-1-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("LastInsertId():", id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("new", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("row affect:", affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	fmt.Println("SELECT * FROM userinfo")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo")
	checkErr(err)

	res, err = stmt.Exec()
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println("row affect:", affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
