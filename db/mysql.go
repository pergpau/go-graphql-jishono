package database

import (
	//"upper.io/db.v3/mysql"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pergpau/go-graphql-jishono/config"
	"log"
)

var Db *sql.DB

func InitDB() {

	db, err := sql.Open("mysql",config.DBConnectionString())
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
 		log.Panic(err)
	}
	Db = db

}