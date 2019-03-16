package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver for
)

// DB is the global database type
var DB *sql.DB

// SMap is is map where the id and value are strings
type SMap map[string]string

// NameAndContent is one database table
type NameAndContent struct {
	Name   string
	Fields SMap
}

// Init sets up the database
func Init() error {
	sqlString := os.Getenv("sqlConnectionString")

	if len(sqlString) == 0 {
		sqlString = "root:markdepro@/a-not-so-secure-site"
	}

	db, err := sql.Open("mysql", sqlString)
	DB = db

	if err != nil {
		return err
	}

	return RefreshDB()
}

// RefreshDB creates the needed tables and fills them with some useless data
// If there are tables it removes them and does the above
func RefreshDB() error {
	tables := []NameAndContent{
		{
			Name: "users",
			Fields: SMap{
				"username": "varchar(255)",
				"password": "varchar(255)",
			},
		},
		{
			Name: "posts",
			Fields: SMap{
				"title":   "varchar(255)",
				"content": "varchar(5000)",
				"userID":  "int",
			},
		},
	}

	for _, table := range tables {
		DB.Exec(`DROP TABLE ` + "`" + table.Name + "`")

		fields := ""
		for fieldName, fieldType := range table.Fields {
			fields = fields + ",\n" + fieldName + " " + fieldType
		}

		query := `CREATE TABLE ` + table.Name + ` (ID int NOT NULL PRIMARY KEY AUTO_INCREMENT` + fields + `);`
		_, err := DB.Exec(query)
		if err != nil {
			return err
		}
	}

	err := SeedUsers()
	if err != nil {
		return err
	}

	err = SeedPosts()
	if err != nil {
		return err
	}

	return nil
}
