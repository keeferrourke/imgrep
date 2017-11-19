package storage

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Remove("imgrep-test.db")
	InitDB("imgrep-test.db")
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	stmt, err := db.Prepare("insert into images (filename, keywords) values (?, ?)")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	filename := "foo.jpeg"
	keywords := "foo,bar"

	_, err = stmt.Exec(filename, keywords)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	results, err := Get("Foo", true) // ignore case specifiers
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if len(results) == 0 {
		t.Fatalf("no results were found")
	}
	if results[0] != filename {
		t.Errorf("%s was not found", filename)
	}

	results, err = Get("foo", false) // respect case specifiers
	if err != nil {
		t.Fatalf("error %v", err)
	}
	if len(results) == 0 {
		t.Fatalf("no results were found")
	}
	if results[0] != filename {
		t.Errorf("%s was not found", filename)
	}
}

func TestInsert(t *testing.T) {
	filename := "bar.jpeg"
	keyword1 := "bar"
	keyword2 := "foobar"

	err := Insert(filename, keyword1, keyword2)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	// perform get to verify insert
	rows, err := db.Query(`select * from images where keywords like ?`, fmt.Sprintf("%%%s%%", keyword1))
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	defer rows.Close()

	var fname string
	var keywords string

	for rows.Next() {
		err := rows.Scan(&fname, &keywords)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
	}

	if fname != filename {
		t.Errorf("Incorrect filename: %v", fname)
	}

	if keywords != fmt.Sprintf("%s,%s", keyword1, keyword2) {
		t.Errorf("Incorrect keyword: %v", keywords)
	}
}

func TestUpdate(t *testing.T) {
	stmt, err := db.Prepare("insert into images (filename, keywords) values (?, ?)")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	filename := "widget.jpeg"
	keywords := "widget"
	newkw := "bangbang"

	_, err = stmt.Exec(filename, keywords)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	err = Update(filename, newkw)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	rows, err := db.Query(`select * from images where filename=?`, filename)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	defer rows.Close()

	var fname string
	var kws string

	for rows.Next() {
		err := rows.Scan(&fname, &kws)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
	}

	if fname != filename {
		t.Errorf("Incorrect filename: %v", fname)
	}

	if kws != newkw {
		t.Errorf("Incorrect keyword: %v", kws)
	}

}
