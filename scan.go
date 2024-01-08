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
	"strconv"
	"strings"
	"time"
)

func FileType(path string) int {
	// 0: undefined/unset/indifferent
	// 1: photo
	// 2: video
	ext := strings.ToLower(filepath.Ext(path))
	//	fmt.Println(ext)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return 1
	case ".mp4", ".avi", ".ts", ".flv", ".wmv", ".mov", ".mpg", ".mpeg", ".mkv", ".webm":
		return 2
	}
	return 0
}

func FileApp(ft int) string {
	if ft == 1 {
		return App.config.Viewer
	} else if ft == 2 {
		return App.config.Player
	}
	return ""
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
	root := "/home/q/k2" //
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

type Dup struct {
	Sha2  string `db:"sha2"`
	Count int    `db:"COUNT(*)"`
}

func RemoveDups() {
	query := "SELECT sha2, COUNT(*) FROM files GROUP BY sha2 HAVING COUNT(*) > 1"
	var dups []Dup
	// select all dups
	err := App.db.Select(&dups, query)
	if err != nil {
		fmt.Println("RemoveDups get dups error: ", err)
		return
	}
	// loop
	var counter int = 1
	for _, d := range dups {
		// select dups
		var media []Media
		q := fmt.Sprintf("SELECT * FROM files WHERE sha2 = X'%x'", d.Sha2)
		e := App.db.Select(&media, q)
		if e != nil {
			fmt.Println("RemoveDups get media error", e)
			return
		}
		for k := range media {
			media[k].Name = filepath.Base(media[k].Path)
		}

		// remove
		var priority = func(s string) int {
			if strings.HasPrefix(s, "h:\\k2\\kawaii") {
				return 5
			}
			if strings.HasPrefix(s, "h:\\k2\\_TWITTER_MY") { // lowest
				return 4
			}
			if strings.HasPrefix(s, "h:\\k2\\_TWITTER") {
				return 3
			}
			if strings.HasPrefix(s, "h:\\k2\\_twi-t") {
				return 2
			}
			if strings.HasPrefix(s, "h:\\k2\\_twi") { // highest
				return 1
			}
			return 6
		}

		var delete = func(index int) {
			fmt.Println("Delete: ", media[index].Path)

			// delete from filesystem
			err := os.Rename(media[index].Path, "h:/k0/"+strconv.Itoa(counter)+"@"+media[index].Name)
			if err != nil {
				fmt.Println(err)
				return
			}
			counter = counter + 1

			// delete from database
			query := `DELETE FROM files WHERE id = ?`
			result, err := App.db.Exec(query, media[index].Id)
			if err != nil {
				log.Fatalln(err)
			}
			count, err := result.RowsAffected()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(count) // prints the number of rows affected

			// delete from slice
			media[index] = media[len(media)-1]
			media = media[:len(media)-1]
		}
		fmt.Println("Count: ", len(media))
		for len(media) > 1 {
			if media[0].Sha1 == media[1].Sha1 && media[0].Sha2 == media[1].Sha2 && media[0].Size == media[1].Size {
				// compare two files
				p0 := priority(media[0].Path)
				p1 := priority(media[1].Path)
				if p1 > p0 {
					delete(1)
				} else if p0 > p1 {
					delete(0)
				} else {
					s0 := strings.Contains(media[0].Name, " ")
					s1 := strings.Contains(media[1].Name, " ")
					if s0 && !s1 {
						delete(0)
					} else if !s0 && s1 {
						delete(1)
					} else if len(filepath.Base(media[0].Name)) > len(filepath.Base(media[1].Name)) {
						delete(0)
					} else {
						delete(1)
					}
				}
			} else {
				fmt.Println("Strange: not equal:\n  ", media[0].Path, "\n  ", media[1].Path)
			}
		}
	}
}
