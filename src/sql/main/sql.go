package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/testdb?charset=utf8")
	checkErr(err)

	//插入数据
	//	stmt, err := db.Prepare("INSERT INTO `testdb`.`users` (`USERNAME`, `PASSWORD`, `MOBILE`, `EMAIL`) VALUES (?, ?, ?, ?);")
	//	checkErr(err)
	//	res, err := stmt.Exec("goname", "gopass", "18181818", "go@mail.com")
	//	checkErr(err)
	//
	//	id, err := res.LastInsertId()
	//	checkErr(err)
	//
	//	fmt.Println(id)

	//查询数据
	rows, err := db.Query("SELECT * from users")
	checkErr(err)
	for rows.Next() {
		var id int
		var username string
		var password string
		var phone string
		var email string
		rows.Scan(&id, &username, &password, &phone, &email)
		checkErr(err)
		fmt.Print(strconv.Itoa(id) + " ")
		fmt.Print(username + " ")
		fmt.Print(password + " ")
		fmt.Print(phone + " ")
		fmt.Print(email + " ")
		fmt.Println()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
