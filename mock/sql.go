package mock

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func SelectAll(db *sql.DB) *sql.Rows {
	//path := strings.Join([]string{"root:pass1234@tcp(", "localhost:3306", ")/", "ktable2333"}, "")
	//db, err := sql.Open("mysql", path)
	//if err != nil {
	//	fmt.Printf("open db err: %+v\n", err)
	//	return nil
	//}

	rows, err := db.Query("select * from ktable_cluster")
	if err != nil {
		fmt.Printf("select err: %+v\n", err)
		return nil
	}
	// 外部传来的db，就不要close了，有外部进行控制，不要越权
	//defer db.Close()

	return rows
}

func UpdateAll(db *sql.DB) sql.Result {
	//path := strings.Join([]string{"root:pass1234@tcp(", "localhost:3306", ")/", "ktable"}, "")
	//db, err := sql.Open("mysql", path)
	//if err != nil {
	//	fmt.Printf("open db err: %+v\n", err)
	//	return nil
	//}

	ret, err := db.Exec("update ktable_cluster set id=123 where id=9999999")
	if err != nil {
		fmt.Printf("update err: %+v\n", err)
		return nil
	}
	//defer db.Close()

	return ret
}