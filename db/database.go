package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"mutlicontainer/constant"
)

type Database interface {
	GetDBinstance() *sql.DB
}

type database struct {
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort int
}

func NewDatabase() database {
	return database{
		DBUser: constant.DBUser,
		DBPassword: constant.DBPassword,
		DBName: constant.DBName,
		DBHost: constant.DBHost,
		DBPort: constant.DBPort,
	}
}

//ConnectDatabase
func (d database) GetDBinstance() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	d.DBHost,d.DBPort,d.DBUser,d.DBPassword,d.DBName)
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