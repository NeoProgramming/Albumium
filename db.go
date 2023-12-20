package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"path/filepath"
	"strings"
)

type Media struct {
	Id    int    `db:"id"`
	Path  string `db:"path"`
	Size  int64  `db:"size"`
	CTime int64  `db:"ctime"`
	MTime int64  `db:"mtime"`
	Sha1  string `db:"sha1"`
	Sha2  string `db:"sha2"`
	FType int
}

const SQLITE_SCHEMA_Files string = `CREATE TABLE IF NOT EXISTS "files" (
	"id"	integer PRIMARY KEY AUTOINCREMENT,
	"path"	varchar(255) DEFAULT "",
	"size"	bigint DEFAULT 0,
	"ctime"	bigint DEFAULT 0,
	"mtime"	bigint DEFAULT 0,
	"sha1"	varchar(255) DEFAULT "",
	"sha2"	varchar(255) DEFAULT ""
)`

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

func getMediaCount(db *sqlx.DB) int {
	var count int
	err := db.Get(&count, "SELECT count(*) FROM files")
	if err != nil {
		return 0
	}
	return count
}

func getMedia(db *sqlx.DB, page int, pageSize int, search string, filters string, order string, desc bool) []Media {
	var media []Media
	query := fmt.Sprintf("SELECT path, size, mtime FROM files")

	if search != "" {
		query += fmt.Sprintf(" WHERE path LIKE '%%%s%%'", search)
	}
	//	if filters != "" {
	//		m, im := decodeFilterMasks(filters)
	//		if search != "" {
	//			query += " AND"
	//		} else {
	//			query += " WHERE"
	//		}
	//		query += fmt.Sprintf(" attrs & %d = %d AND attrs & %d = 0", m, m, im)
	//	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}
	query += ";"

	fmt.Println("getMedia: ", query)
	err := db.Select(&media, query)
	if err != nil {
		fmt.Println("getMedia error", err)
		return nil
	}
	for i := range media {

		media[i].Path = filepath.ToSlash(media[i].Path)
		media[i].Path = strings.Replace(media[i].Path, App.config.BasePath, "", 1)
		media[i].FType = FileType(media[i].Path)

	//	fmt.Println(media[i].Path, "  == ", media[i].FType)
	}
	return media
}
