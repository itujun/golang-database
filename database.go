package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// parseTime = true -> mengubah tipe data dari database ke format time.Time milik golang
	db, err := sql.Open("mysql", "root:@/learn-golang-database?parseTime=true")
	if err != nil {
		panic(err)
	}

	// Setting db pooling
	db.SetMaxIdleConns(10) // set minimal 10 koneksi
	db.SetMaxOpenConns(100) // set maksimal 100 koneksi
	db.SetConnMaxIdleTime(5 * time.Minute) // set koneksi yang sudah tidak digunakan selama 5 menit akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // set koneksi boleh digunakan selama 60 menit

	return db
}