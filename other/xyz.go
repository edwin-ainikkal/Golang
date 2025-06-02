package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

type Data struct {
	id   int
	name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	result, err := db.Exec("insert into data values(4, 'xyz')")
	checkError(err)
	lastinsert, err := result.LastInsertId()
	checkError(err)
	fmt.Println("Id :", lastinsert)
	rowsaffected, err := result.RowsAffected()
	checkError(err)
	fmt.Println("Rows affected :", rowsaffected)

	rows, err := db.Query("Select * from data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}

}
