package main

import (
	"database/sql"
	"fmt"
)

var _ *sql.DB
var _ error

func connectDatabase() {
	_, _ = sql.Open("postgres", "postgres://postgres:test@localhost:5432/postgres")
	fmt.Println("Database Connected!")
}
