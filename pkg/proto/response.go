package proto

//go:generate msgp

type Response struct {
	Id           uint32 `msg:"i"`
	RowsAffected int64  `msg:"ra"`
	LastInsertId int64  `msg:"li"`
	Error        string `msg:"e"`
}
