package proto

//go:generate msgp

type RequestError struct {
	Message   string `msg:"msg"`
	ErrorCode int    `msg:"error_code"`
}

type Response struct {
	Id     uint32        `msg:"id"`
	Result []interface{} `msg:"result"`
	Error  RequestError  `msg:"error"`
}
