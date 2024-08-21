package facade

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nesterow/dal/pkg/adapter"
	"github.com/nesterow/dal/pkg/handler"
)

type SQLiteServer struct {
	Cwd  string
	Port string
}

/*
Init initializes the SQLiteServer struct with the required environment variables.
- `SQLITE_DIRECTORY` is the directory where the SQLite database is stored.
- `SQLITE_PORT` is the port on which the server listens.
*/
func (s *SQLiteServer) Init() {
	s.Cwd = os.Getenv("SQLITE_DIRECTORY")
	if s.Cwd == "" {
		panic("env variable `SQLITE_DIRECTORY` is not set")
	}
	os.MkdirAll(s.Cwd, os.ModePerm)
	os.Chdir(s.Cwd)
	s.Port = os.Getenv("SQLITE_PORT")
	if s.Port == "" {
		s.Port = "8118"
	}
}

/*
GetAdapter returns a DBAdapter struct with the SQLite3 dialect registered.
- The `SQLITE_PRAGMAS` environment variable is expected to be a semicolon separated list of PRAGMA statements.
*/
func (s *SQLiteServer) GetAdapter() adapter.DBAdapter {
	adapter.RegisterDialect("sqlite3", adapter.CommonDialect{})
	db := adapter.DBAdapter{
		Type: "sqlite3",
	}
	db.AfterOpen("PRAGMA journal_mode=WAL")
	for _, pragma := range strings.Split(os.Getenv("SQLITE_PRAGMAS"), ";") {
		if pragma == "" {
			continue
		}
		db.AfterOpen(pragma)
	}
	return db
}

/*
GetHanfler returns a http.Handler configured for the SQLiteServer.
*/
func (s *SQLiteServer) GetHandler() http.Handler {
	return handler.QueryHandler(s.GetAdapter())
}

/*
Serve starts the basic server on the configured port.
Use `GetHandler` to get a handler for a custom server.
*/
func (s *SQLiteServer) Serve() {
	s.Init()
	log.Println("Starting server on port " + s.Port)
	log.Println("Using directory: " + s.Cwd)
	err := http.ListenAndServe(":"+s.Port, s.GetHandler())
	if err != nil {
		panic(err)
	}
}
