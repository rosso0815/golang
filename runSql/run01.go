package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Printf("exit 1")
	os.Exit(1)

	db, err := sql.Open("mysql", "play:play@/play?charset=utf8")
	checkErr(err)

	res, err := db.Exec("drop table userinfo if exists")
	checkErr(err)

	// create table
	res, err = db.Exec("create table userinfo ( username varchar(20) )")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?")
	checkErr(err)

	res, err = stmt.Exec("astaxie")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
