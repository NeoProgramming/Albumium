package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Application struct {
	db *sqlx.DB
}

var App Application

func InitDatabase() {
	db, err := sqlx.Connect("sqlite3", "./album.db")
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	App.db = db

	if _, err := App.db.Exec(SQLITE_SCHEMA_Files); err != nil {
		log.Fatal("error crate table: ", err)
		panic(err)
	}

	App.db.Exec("PRAGMA journal_mode = WAL")
	App.db.Exec("PRAGMA synchronous = normal")
	App.db.Exec("PRAGMA temp_store = memory")
	App.db.Exec("PRAGMA mmap_size = 30000000000")
	App.db.Exec("PRAGMA case_sensitive_like = false")

	fmt.Println("Database album.db opened")
}

func CloseDatabase() {
	App.db.Close()
}

func Insert(path string, size int64, mtime time.Time) {

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h1 := sha1.New()
	if _, err := io.Copy(h1, f); err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("h1 = %x\n", h1.Sum(nil))
	f.Seek(0, 0)

	h2 := sha256.New()
	if _, err := io.Copy(h2, f); err != nil {
		log.Fatal(err)
	}

	if _, err := App.db.Exec("INSERT into Files(path, size, mtime, sha1, sha2) VALUES (?, ?, ?, ?, ?)",
		path, size, mtime.Unix(), h1.Sum(nil), h2.Sum(nil)); err != nil {
		fmt.Println("insert error: path=", path, " err=", err)
		panic(err)
	}

}

func Scan() {
	// recursive scan
	root := "h:/k2" // "h:/k2/!UNS-IMG/!111" //
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)
		// add path to database
		if !info.IsDir() {
			Insert(path, info.Size(), info.ModTime())
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
	}
}

func main() {
	fmt.Printf("hello")
	InitDatabase()
	Scan()
	CloseDatabase()
	fmt.Printf("good bye")
}
