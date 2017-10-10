package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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

type Result struct {
	Filename string   `json:"filename"`
	Keywords []string `json:"keywords"`
}

func Get(keyword string) ([]string, error) {
	if keyword == "" {
		return nil, errors.New("query: empty query matches all images")
	}
	results := []string{}
	rows, err := db.Query(`select * from images`, fmt.Sprintf("%%%s%%", keyword))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var filename string
		var keywords string

		err := rows.Scan(&filename, &keywords)
		if err != nil {
			return nil, err
		}

		found := false
		for _, kw := range strings.Split(keywords, ",") {
			if strings.Contains(strings.ToLower(kw), strings.ToLower(keyword)) {
				found = true
			}
		}

		if found {
			results = append(results, filename)
		}
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

func Delete(filename string) error {
	stmt, err := db.Prepare("delete from images where filename=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(filename)
	return err
}
