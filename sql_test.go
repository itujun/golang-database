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
	// gunakan perintah Exec/ExecContext untuk perintah Sql yang tidak membutuhkan hasil / return
	_, err := db.ExecContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil disimpan");
}

func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "SELECT id, name FROM customer"
	// gunakan perintah Query/QueryContext untuk perintah Sql yang membutuhkan hasil / return
	rows, err := db.QueryContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "| name:", name)
	}
}