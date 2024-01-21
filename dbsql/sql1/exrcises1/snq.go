package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Albom struct {
	id        int
	fist_name string
	last_name string
	email     string
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "yenisifre"
	dbname   = "DVDrental"
)

var db *sql.DB

func main() {

	users := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err := sql.Open("postgres", users)
	if err != nil {
		log.Fatal()
	}

	dberor := db.Ping()
	if dberor != nil {
		log.Fatal(dberor)
	} else {
		fmt.Println("connect")
	}

}

func select1()
