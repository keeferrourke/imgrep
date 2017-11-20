package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	// go-sqlite3 functions are used everywhere in this file, so a blank import
	// is appropriate
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB initializes an sqlite3 database for imgrep at the specified path
func InitDB(path string) error {
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
		return err
	}
	return nil
}

// Insert inserts a filename-keyword mapping as a row in the database
func Insert(filename string, keywords ...string) error {
	keywordsAppended := strings.Join(keywords, ",")
	stmt, err := db.Prepare("insert into images (filename, keywords) values (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(filename, keywordsAppended)
	return err
}

// Remove removes a row containing the specified filename
func Remove(filename string) error {
	stmt, err := db.Prepare("delete from images where filename = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(filename)
	return err
}

// Lookup verifies that a filename is in the database
func Lookup(filename string) bool {
	if filename == "" {
		return false
	}
	var fn, kw string
	err := db.QueryRow("select * from images where filename = ?", filename).Scan(&fn, &kw)
	if err != nil {
		return false
	}
	return true
}

// Get retrieves names of image files which match the keyword
func Get(keyword string, ignoreCase bool) ([]string, error) {
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
			if ignoreCase {
				if strings.Contains(strings.ToLower(kw), strings.ToLower(keyword)) {
					found = true
				}
			} else {
				if strings.Contains(kw, keyword) {
					found = true
				}
			}
		}

		if found {
			results = append(results, filename)
		}
	}
	return results, nil
}

// Update updates a database entry
func Update(filename string, keywords ...string) error {
	keywordsAppended := strings.Join(keywords, ",")
	stmt, err := db.Prepare("update images set keywords=? where filename=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(keywordsAppended, filename)
	return err
}
