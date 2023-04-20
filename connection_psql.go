package goutils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

/*
ConnectionBDPostgreSQL o antigo nome era: ConexaoDB
*/
func ConnectionBDPostgreSQL(applicationName string) *sql.DB {
	var host, dbname, user, password string
	var port int

	host = Godotenv("host")
	dbname = Godotenv("dbname")
	port, _ = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable application_name=%s", host, port, user, password, dbname, applicationName)

	db, err := sql.Open("postgres", psqlInfo)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		time.Sleep(2 * time.Second)
		return ConnectionBDPostgreSQL(applicationName)
	}

	return db
}

func ConnectionBDPostgreSQLWithSSL() *sql.DB {
	var host, dbname, user, password string
	var port int

	host = Godotenv("host")
	dbname = Godotenv("dbname")
	port, _ = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		time.Sleep(2 * time.Second)
		return ConnectionBDPostgreSQL()
	}

	return db
}

/*
ConnectionBDPostgreSQLORM o antigo nome era: ConexaoBDORM
*/
func ConnectionBDPostgreSQLORM() (DB *gorm.DB) {
	var host, dbname, user, password string
	var port int

	host = Godotenv("host")
	dbname = Godotenv("dbname")
	port, _ = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
