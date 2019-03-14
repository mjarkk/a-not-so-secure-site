package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver for
)

// Init sets up the database
func Init() error {
	sqlString := os.Getenv("sqlConnectionString")

	if len(sqlString) == 0 {
		sqlString = "root:markdepro@/dbname"
	}

	db, err := sql.Open("mysql", sqlString)
	fmt.Println(db, err)
	return err
}
