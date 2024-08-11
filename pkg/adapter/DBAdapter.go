package adapter

import (
	"database/sql"
	"fmt"
	"time"
)

/**
 * DBAdapter
 * Automatically creates connections for each database URL.
 * Executes queries on the specified database.
 * TODO: Closes connections older than 2 minutes.
**/
type DBAdapter struct {
	Type              string
	MaxAttempts       int
	_db               map[string]*sql.DB
	_connect_attempts map[string]int
	_connect_time     map[string]int64
}

func (a *DBAdapter) Open(url string) (*sql.DB, error) {
	if a._db == nil {
		a._db = make(map[string]*sql.DB)
	}
	if a._connect_attempts == nil {
		a._connect_attempts = make(map[string]int)
	}
	if a._connect_time == nil {
		a._connect_time = make(map[string]int64)
	}
	if a.MaxAttempts == 0 {
		a.MaxAttempts = 6
	}
	if _, ok := a._db[url]; !ok {
		a._db[url], _ = sql.Open(a.Type, url)
	} else {
		err := a._db[url].Ping()
		if err != nil {
			a._db[url] = nil
			a._connect_attempts[url]++
			time.Sleep(time.Duration(5) * time.Second)
			if a._connect_attempts[url] > a.MaxAttempts {
				return nil, fmt.Errorf(`failed to connect to "%s"`, url)
			}
			return a.Open(url)
		}
	}
	a._connect_attempts[url] = 0
	a._connect_time[url] = time.Now().Unix()
	return a._db[url], nil
}

func (a *DBAdapter) GetDB(url string) *sql.DB {
	if a._db == nil {
		return nil
	}
	return a._db[url]
}

func (a *DBAdapter) Close() {
	for _, db := range a._db {
		db.Close()
	}
	a._db = nil
	a._connect_attempts = nil
}

func (a *DBAdapter) Query(req Query) (*sql.Rows, error) {
	db, err := a.Open(req.Db)
	if err != nil {
		return nil, err
	}
	sfmt, err := db.Prepare(req.Expression)
	if err != nil {
		return nil, err
	}
	return sfmt.Query(req.Data...)
}

func (a *DBAdapter) Exec(req Query) (sql.Result, error) {
	db, err := a.Open(req.Db)
	if err != nil {
		return nil, err
	}
	sfmt, err := db.Prepare(req.Expression)
	if err != nil {
		return nil, err
	}
	return sfmt.Exec(req.Data...)
}
