package config

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBInit() (db *sql.DB) {

	// dbdriver := "mysql"
	// dbhost := "localhost" //"147.139.139.200"
	// dbport := "3306"         //"6044"
	// dbname := "sesi11"
	// dbuser := "root"              //"remote"
	// dbpass := "" //"kucinglucu"

	//db, err := gorm.Open(dbdriver, dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?timeout=600s&charset=utf8&parseTime=false&loc=Asia%2FJakarta")
	//db, err := gorm.Open("mysql", "root:@/sesi11?charset=utf8")
	db, err := sql.Open("mysql", "root:@/sesi11?charset=utf8")
	// db, err := gorm.Open("mysql", "root:P@ssw0rd@tcp(172.18.133.135:33306)/ifs?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
