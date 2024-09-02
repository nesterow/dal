package facade

import (
	"database/sql"
	"reflect"

	"l12.xyz/x/dal/pkg/adapter"
	"l12.xyz/x/dal/pkg/proto"
)

var db adapter.DBAdapter

func InitSQLite(pragmas []string) {
	if db.Type == "" {
		adapter.RegisterDialect("sqlite3", adapter.CommonDialect{})
		db = adapter.DBAdapter{
			Type: "sqlite3",
		}
		db.AfterOpen("PRAGMA journal_mode=WAL")
	}
	for _, pragma := range pragmas {
		if pragma == "" {
			continue
		}
		db.AfterOpen(pragma)
	}
}

type RowsIter struct {
	Result  []byte
	Columns []string
	rows    *sql.Rows
}

func (r *RowsIter) Exec(input []byte) {
	InitSQLite([]string{})
	req := proto.Request{}
	_, e := req.UnmarshalMsg(input)
	query, err := req.Parse(adapter.GetDialect(db.Type))
	if err != nil || e != nil {
		res := proto.Response{
			Error: "failed to unmarshal request",
		}
		r.Result, _ = res.MarshalMsg(nil)
		return
	}
	if query.Exec {
		result, err := db.Exec(query)
		if err != nil {
			res := proto.Response{
				Error: err.Error(),
			}
			r.Result, _ = res.MarshalMsg(nil)
			return
		}
		ra, _ := result.RowsAffected()
		la, _ := result.LastInsertId()
		res := proto.Response{
			Id:           0,
			RowsAffected: ra,
			LastInsertId: la,
		}
		r.Result, _ = res.MarshalMsg(nil)
		return
	}
	rows, err := db.Query(query)
	if err != nil {
		res := proto.Response{
			Error: err.Error(),
		}
		r.Result, _ = res.MarshalMsg(nil)
		return
	}
	r.rows = rows
}

func (r *RowsIter) Close() {
	if r.rows == nil {
		return
	}
	r.rows.Close()
}

func (r *RowsIter) Next() []byte {
	columns, _ := r.rows.Columns()
	types, _ := r.rows.ColumnTypes()
	if r.Columns == nil {
		r.Columns = columns
		cols, _ := proto.MarshalRow(columns)
		return cols
	}
	data := make([]interface{}, len(columns))
	r.rows.Next()
	for i := range data {
		typ := reflect.New(types[i].ScanType()).Interface()
		data[i] = &typ
	}
	r.rows.Scan(data...)
	cols, _ := proto.MarshalRow(data)
	return cols
}
