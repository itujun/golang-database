package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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


/**
* ? MAPPING TIPE DATA
** |======== TIPE DATA DATABASE ======|===== TIPE DATA GOLANG =====|
*  | VARCHAR, CHAR					 					 | string               		 |
*  | INT, BIGINT						 					 | int32, int64						 	 |
*  | FLOAT, DOUBLE					 					 | float32, float64					 |
*  | BOOLEAN                					 | bool	               			 |
*  | DATE, DATETIME, TIME, TIMESTAMP	 | time.Time              	 |
*/

/**
* ? TIPE DATA NULLABLE
** |======== TIPE DATA GOLANG ======|===== TIPE DATA NULLABLE =====|
*  | string					 					 			| database/sql.NullString			|
*  | bool						 					 			| database/sql.NullBool				|
*  | float64					 					 		| database/sql.NullFloat64		|
*  | int32                					| database/sql.NullInt32			|
*  | int64 						             	| database/sql.NullInt64			|
*  | time.Time              	 			|	database/sql.NullTime				|
*/

func TestQuerySqlComplex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	// gunakan perintah Query/QueryContext untuk perintah Sql yang membutuhkan hasil / return
	rows, err := db.QueryContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime 
		var created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "| name:", name, "| email:", email, "| balance:", balance, "| rating:", rating, "| birth_date:", birth_date, "| married:", married, "| created_at:", created_at, "Email Not Null?", email.Valid, "| Birth Date Null?", birth_date.Valid);
		fmt.Println("======================");
	}
}