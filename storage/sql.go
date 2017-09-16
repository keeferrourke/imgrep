package storage

import (
	"database/sql"
	"fmt"
	"log"
    "os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const DBPATH string = "storage" + string(os.PathSeparator) + "imgrep.db"

func InitDB(path string) {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
    create table if not exists images (
        filename text not null primary key,
        keywords text
    );
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func Insert(filename string, keywords ...string) error {
	keywordsAppended := strings.Join(keywords, ",")
	stmt, err := db.Prepare("insert into images (filename, keywords) values (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(filename, keywordsAppended)
	return err
}

func Get(keyword string) ([]string, error) {
	results := []string{}
	rows, err := db.Query(`select filename from images where keywords like ?`, fmt.Sprintf("%%%s%%", keyword))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var filename string

		err := rows.Scan(&filename)
		if err != nil {
			return nil, err
		}

		results = append(results, filename)
	}
	return results, nil
}

func Update(filename string, keywords ...string) error {
	keywordsAppended := strings.Join(keywords, ",")
	stmt, err := db.Prepare("update images set keywords=? where filename=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(keywordsAppended, filename)
	return err
}
