package main

import (
	"log"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//
//var db *sql.DB
//var err error
//
//func init() {
//	db, err = sql.Open("mysql", "root:@/School")
//	if err != nil {
//		panic(err)
//	}
//	err = db.Ping()
//	if err != nil {
//		log.Println("DB Connection is Failed")
//		panic(err)
//	}
//}

//func main() {
//	log.Println("DB IS READY")
//	defer db.Close()
//
//	t := time.Now()
//	rows, err := db.Query(`SELECT string2 FROM connect_pool_test LIMIT ?`, 100)
//	if err != nil {
//		panic(err)
//	}
//	log.Println("used time:", time.Now().Sub(t))
//	for i := 0; rows.Next(); i++ {
//		var s sql.NullString
//		err = rows.Scan(&s)
//		if err != nil {
//			log.Println("scan row error =>", err)
//		}
//		if s.Valid {
//			// use s.String
//			fmt.Println("string2[", i, "] =>", s.String)
//		} else {
//			// NULL value
//		}
//	}
//
//}

var id int = 3

func main() {
	//log.Println("DB is Ready.")
	//defer db.Close()

	// insert
	//var sqlInsert bytes.Buffer
	//argsInsert := make([]interface{}, 0, 20)
	//
	//sqlInsert.WriteString("INSERT INTO student (")
	//sqlInsert.WriteString("id,name,birth,department,address) VALUE")
	//sqlInsert.WriteString("(?,?,?,?,?)")
	//argsInsert = append(argsInsert, id, "testName", "1995", "", "")
	//
	//rslt, err := db.Exec(sqlInsert.String(), argsInsert...)
	//if err != nil {
	//	log.Fatal("INSERT False, error:", err)
	//}
	//log.Println("INSERT Success, rslt:", rslt)

	// select
	//var sqlSelect bytes.Buffer
	//var ID int32
	//var Name string
	//var Sex string
	//var Birth string
	//var Department string
	//var Address string
	//var isSelected int32
	//
	//sqlSelect.WriteString("SELECT `id`, `name`, `sex`, `birth`,`department`,`address`, `is_selected` FROM student WHERE `id`=?")
	//stmtOut, err := db.Prepare(sqlSelect.String())
	//defer stmtOut.Close()
	//err = stmtOut.QueryRow(id).Scan(&ID, &Name, &Sex, &Birth, &Department, &Address, &isSelected)
	//if err != nil {
	//	log.Fatal("SELECT False, error:", err)
	//}
	//log.Println("SELECT Success, rslt:", ID, Name, Sex, Birth, Department, Address, isSelected)

	numInt, _ := strconv.Atoi(nil)
	//if err != nil {
	//	log.Fatal("strconv error:", err)
	//}
	log.Println(numInt)

}
