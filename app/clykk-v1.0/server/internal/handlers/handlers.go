package handlers

import (
	"github.com/clykk-user-atif/db"
)

type Handler struct {
	dbStr string

	rDB *db.ReaderDB
	wDB *db.WriterDB
}

func NewHandler(dbStr string) (Handler, error) {
	rDB, err := db.Reader(dbStr)
	if err != nil {
		return Handler{}, err
	}
	wDB, err := db.Writer(dbStr)
	if err != nil {
		return Handler{}, err
	}
	return Handler{
		dbStr: dbStr,
		rDB:   rDB,
		wDB:   wDB,
	}, nil
}
