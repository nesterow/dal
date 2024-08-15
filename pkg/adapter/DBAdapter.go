package adapter

import (
	"database/sql"
	"fmt"
	"time"
)

/*
DBAdapter
Automatically creates connections for each database URL.
Executes queries on the specified database.
Closes connections older than ConnectionLiveTime
*/
type DBAdapter struct {
	Type               string
	MaxAttempts        int
	ConnectionLiveTime int
	dbs                *DBMap
}

type DBMap struct {
	Connections        map[string]*sql.DB
	ConnectionAttempts map[string]int
	ConnectionTime     map[string]int64
}

func (a *DBAdapter) Open(url string) (*sql.DB, error) {
	defer a.CleanUp()

	if a.MaxAttempts == 0 {
		a.MaxAttempts = 6
	}

	if a.ConnectionLiveTime == 0 {
		a.ConnectionLiveTime = 120
	}

	if a.dbs == nil {
		a.dbs = &DBMap{
			Connections:        make(map[string]*sql.DB),
			ConnectionAttempts: make(map[string]int),
			ConnectionTime:     make(map[string]int64),
		}
	}

	connections := a.dbs.Connections
	attempts := a.dbs.ConnectionAttempts
	lastHits := a.dbs.ConnectionTime

	lastHits[url] = time.Now().Unix()
	if _, ok := connections[url]; !ok {
		connections[url], _ = sql.Open(a.Type, url)
	} else {
		err := connections[url].Ping()
		if err != nil {
			connections[url] = nil
			attempts[url]++
			time.Sleep(time.Duration(5) * time.Second)
			if attempts[url] > a.MaxAttempts {
				return nil, fmt.Errorf(
					`failed to connect to "%s", after %v attempts`,
					url,
					a.MaxAttempts,
				)
			}
			return a.Open(url)
		}
	}

	attempts[url] = 0
	return connections[url], nil
}

func (a *DBAdapter) GetDB(url string) *sql.DB {
	if a.dbs == nil {
		return nil
	}
	return a.dbs.Connections[url]
}

func (a *DBAdapter) Close() {
	for url, db := range a.dbs.Connections {
		db.Close()
		delete(a.dbs.Connections, url)
		delete(a.dbs.ConnectionAttempts, url)
		delete(a.dbs.ConnectionTime, url)
	}
	a.dbs = nil
}

func (a *DBAdapter) CleanUp() {
	if a.dbs == nil {
		return
	}
	lastHits := a.dbs.ConnectionTime
	liveTime := a.ConnectionLiveTime
	for url, db := range a.dbs.Connections {
		if time.Now().Unix()-lastHits[url] > int64(liveTime) {
			db.Close()
			delete(a.dbs.Connections, url)
			delete(a.dbs.ConnectionAttempts, url)
			delete(a.dbs.ConnectionTime, url)
		}
	}
}

func (a *DBAdapter) Query(req Query) (*sql.Rows, error) {
	db, err := a.Open(req.Db)
	if err != nil {
		return nil, err
	}
	if req.Transaction {
		tx, _ := db.Begin()
		rows, err := tx.Query(req.Expression, req.Data...)
		tx.Commit()
		return rows, err
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
	if req.Transaction {
		tx, _ := db.Begin()
		result, err := tx.Exec(req.Expression, req.Data...)
		tx.Commit()
		return result, err
	}
	sfmt, err := db.Prepare(req.Expression)
	if err != nil {
		return nil, err
	}
	return sfmt.Exec(req.Data...)
}
