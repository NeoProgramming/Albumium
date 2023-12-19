package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FileType(path string) int {
	ext := strings.ToLower(filepath.Ext(path))
	fmt.Println(ext)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return 0
	case ".mp4", ".avi", ".ts", ".flv", ".wmv", ".mov", ".mpg", ".mpeg", ".mkv", ".webm":
		return 1
	}
	return -1
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
	root := "/home/q/k2" // "h:/k2/!UNS-IMG/!111" //
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
