package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB sql.DB

func InitDB(path string) {
	DB, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	sqlStmt := `
    create table images (
        filename text not null primary key,
        keywords text,
    );
    `
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

func (DB *sql.DB) Insert(keywords ...string) {

}

func (DB *sql.DB) Get(keyword string) {

}
