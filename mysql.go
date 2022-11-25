package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mysql(host string, port int, user, password string, database string) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)
	fmt.Printf("dataSource <%s>\n", s)
	db, err := sql.Open("mysql", s)
	if err != nil {
		fmt.Println("mysql failed", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("select 1 from mysql.user")
	if err != nil {
		fmt.Println("mysql failed", err)
		return
	}
	fmt.Println("mysql success")
}
