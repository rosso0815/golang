package main

import (
	//"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("@@@ init")
	/*
		db, err := sql.Open("mysql", "reader:reader@/employees?charset=utf8")
		checkErr(err)
		rows, err := db.Query("SELECT emp_no,first_name,last_name FROM employees")
		checkErr(err)
		for rows.Next() {
			var empNo int
			var firstName string
			var lastName string
			err = rows.Scan(&empNo, &firstName, &lastName)
			checkErr(err)
			fmt.Printf("emp_no=%d,first_name=%s,last_name=%s\n", empNo, firstName, lastName)
		}*/
}

func Abs() {
	fmt.Println("@@@ Abs")
}

func Sum(x int, y int) int {
	return x + y
}

func TestAbs(t *testing.T) {
	fmt.Println("@@@ TestAbs start")
	// 	// query
	// 	rows, err := db.Query("SELECT emp_no,first_name,last_name FROM employees")
	// 	checkErr(err)
	fmt.Println("@@@ TestAbs finish")
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("found error")
		panic(err)
	}
}

// // go run main.go -user pfistera -conn drone:drone@/drone

// func main() {

// 	// TODO check if connect to NAS

// 	db, err := sql.Open("mysql", "reader:reader@/employees?charset=utf8")
// 	checkErr(err)

// 	// insert
// 	//	stmt, err := db.Prepare("INSERT book SET author=?,title=?,url=?")
// 	//checkErr(err)

// 	//res, err := stmt.Exec("author1", "title01", "http://www.google.ch")
// 	//checkErr(err)

// 	//id, err := res.LastInsertId()
// 	//checkErr(err)

// 	//fmt.Println(id)
// 	// update
// 	//stmt, err = db.Prepare("update book set username=? where uid=?")
// 	//checkErr(err)

// 	//res, err = stmt.Exec("astaxieupdate", id)
// 	//checkErr(err)

// 	//	affect, err := res.RowsAffected()
// 	//	checkErr(err)

// 	//	fmt.Println(affect)

// 	// query
// 	rows, err := db.Query("SELECT emp_no,first_name,last_name FROM employees")
// 	checkErr(err)

// 	for rows.Next() {
// 		var emp_no int
// 		var first_name string
// 		var last_name string
// 		err = rows.Scan(&emp_no, &first_name, &last_name)
// 		checkErr(err)
// 		fmt.Printf("emp_no=%d,first_name=%s,last_name=%s\n", emp_no, first_name, last_name)
// 	}

// 	// delete
// 	//stmt, err = db.Prepare("delete from userinfo where uid=?")
// 	//checkErr(err)

// 	//res, err = stmt.Exec(id)
// 	//checkErr(err)

// 	//affect, err = res.RowsAffected()
// 	//checkErr(err)

// 	//fmt.Println(affect)

// 	db.Close()
// }
