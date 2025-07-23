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

	/**
	 * ? Pengaturan Database Pooling
	 * * (DB) SetMaxIdleConns(number) / saat lagi ngganggur
	 * -> Pengaturan berapa jumlah koneksi minimal yang dibuat
	 * * (DB) SetMaxOpenConns(number)
	 * -> Pengaturan berapa jumlah koneksi maksimal yang dibuat
	 * * (DB) SetConnMaxIdleTime(duration) / saat lagi ngganggur
	 * -> Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
	 * * (DB) SetConnMaxLifetime(duration)
	 * -> Pengaturan berapa lama koneksi boleh digunakan
	 */

}