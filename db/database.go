package db

import (
	"database/sql"
	"fmt"

	"mutlicontainer/configuration"
	"log"

	_ "github.com/lib/pq"
)

// type Database interface {
// 	Instance() *sql.DB
// }

//Config
type database struct {
	config configuration.ProductConfig
}

//NewConfig
func NewDatabase(config configuration.ProductConfig) database {
	return database{
		config: config,
	}
}

//ConnectDatabase
func GetDBinstance(d *database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.config.GetDBHost(), d.config.GetDBPort(), d.config.GetDBUser(), d.config.GetDBPassword(), d.config.GetDBName())
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db is used=========")
	return conn
}


