package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "INSERT INTO customer(id, name) VALUES('lev', 'Lev')"
	_, err := db.ExecContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil disimpan");
}