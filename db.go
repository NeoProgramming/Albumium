package main

type File struct {
	Id    int    `db:"id"`
	Path  string `db:"path"`
	Size  int64  `db:"size"`
	CTime int64  `db:"ctime"`
	MTime int64  `db:"mtime"`
	Sha1  string `db:"sha1"`
	Sha2  string `db:"sha2"`
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
