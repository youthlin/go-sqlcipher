package main

import (
	"database/sql"

	_ "github.com/youthlin/go-sqlcipher"
)

func main() {
	for _, driver := range sql.Drivers() {
		println(driver)
	}
}
