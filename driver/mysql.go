package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectToSQL() *sql.DB {
	// get a database handle
	db, err := sql.Open("mysql", "root:Anur@120799@tcp(0.0.0.0:3306)/blogs")
	if err != nil {
		log.Println(err)
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
	}

	log.Println("Connected!")

	return db
}