package belajar_golang_database1

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/pzn_belajar_golang_database1")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// gunakan DB
}
