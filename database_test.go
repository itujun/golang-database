package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/learn-golang-database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}