package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:@/mysql_benchmark")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("DB Connection is Failed")
		panic(err)
	}
}

func main() {
	log.Println("DB IS READY")
	defer db.Close()

	t := time.Now()
	rows, err := db.Query(`SELECT string2 FROM connect_pool_test LIMIT ?`, 100)
	if err != nil {
		panic(err)
	}
	log.Println("used time:", time.Now().Sub(t))
	for i := 0; rows.Next(); i++ {
		var s sql.NullString
		err = rows.Scan(&s)
		if err != nil {
			log.Println("scan row error =>", err)
		}
		if s.Valid {
			// use s.String
			fmt.Println("string2[", i, "] =>", s.String)
		} else {
			// NULL value
		}
	}

}
