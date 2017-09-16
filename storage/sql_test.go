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
	keywords := "yolo,yolobolo"

	_, err = stmt.Exec(filename, keywords)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	results, err := Get("yolo")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if len(results) == 0 {
		t.Fatalf("no results were found")
	}

	if results[0] != filename {
		t.Error("f.jpeg was not found")
	}
}

func TestInsert(t *testing.T) {
	filename := "foo2.jpeg"
	keyword1 := "yolo"
	keyword2 := "yolobolo"

	err := Insert(filename, keyword1, keyword2)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

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
