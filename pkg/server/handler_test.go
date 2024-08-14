package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"l12.xyz/dal/adapter"
)

func TestQueryHandler(t *testing.T) {
	body := []byte(`{
		"something": "wrong",
	}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := QueryHandler(adapter.DBAdapter{
		Type: "sqlite3",
	})
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Code == 400)
}
