package sqlite

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	instance *sql.DB
	once     sync.Once
)

func Connect(dsName string) *sql.DB {
	once.Do(func() {
		db, err := sql.Open("sqlite3", dsName)
		if err != nil {
			panic(err)
		}
		instance = db
	})

	return instance
}
