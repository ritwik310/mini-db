package server_test

import (
	"testing"

	"github.com/ritsource/mini-db/server"
	"github.com/ritsource/mini-db/src"
)

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestHandleMsg(t *testing.T) {
	store := src.Store{Persist: false}
	store.Map = make(map[string]interface{})

	// Testing SET-Command hanlding
	bs1 := server.HandleMsg(&store, []byte("SET\r\nkey1\r\n+OK\r\n"))
	d1, err := src.UnmarshalData(bs1)
	handleErr(t, err)

	if d1["error"] != nil {
		t.Error("d1[\"error\"] != nil")
	}

	// Testing GET-Command hanlding
	bs2 := server.HandleMsg(&store, []byte("GET\r\nkey1\r\n"))
	d2, err := src.UnmarshalData(bs2)
	handleErr(t, err)

	if d2["error"] != nil {
		t.Error("d2[\"error\"] != nil")
	} else if d2["data"] != "OK" {
		t.Error("d2[\"data\"] != \"OK\"")
	}

	// Testing DELETE-Command
	bs3 := server.HandleMsg(&store, []byte("DELETE\r\nkey1\r\n"))
	d3, err := src.UnmarshalData(bs3)
	handleErr(t, err)

	if d3["error"] != nil {
		t.Error("d3[\"error\"] != nil")
	}

	// Testing SET-Cmd for non existing key
	bs4 := server.HandleMsg(&store, []byte("GET\r\nkey1\r\n"))
	d4, err := src.UnmarshalData(bs4)
	handleErr(t, err)

	if d4["error"] == nil {
		t.Error("d4[\"error\"] == nil", d4["error"], "or unable to delete file")
	}

	// Testing SET-Command where value contains string
	bs5 := server.HandleMsg(&store, []byte("SET\r\nkey1\r\n+Ritwik\r\n+Saha\r\n"))
	d5, err := src.UnmarshalData(bs5)
	handleErr(t, err)

	if d5["error"] != nil {
		t.Error("d1[\"error\"] != nil")
	}
}
