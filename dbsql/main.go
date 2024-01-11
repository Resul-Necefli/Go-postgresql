/*
GO (golang)  ile  database  baglanma
temel  prosesler

*/

package main

import (
	"database/sql"
	"fmt"
	"log"

	Insertgo "github.com/Resul-Necefli/Go-postgresql/insert"
	Sqll "github.com/Resul-Necefli/Go-postgresql/sql"
	_ "github.com/lib/pq"
)

/*
database  melumatlarimizi  bu  sekilde qeyd
etdiy cunki istfade zamani  daha  az sefe  yolvermis olariq

*/

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "yenisifre"
	dbname   = "DVDrental"
	staff    = "staff"
)

var db *sql.DB //  database  paketimizn butun funksialarini bu  deyiskende  tanimladiq  ve  qlobal  seviyyede

// database  baglantini yardaq
func InitDB() {
	constString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", constString)
	if err != nil {
		log.Fatal(err)
	}
}

func Updatesql() {
	_, err := db.Exec("UPDATE "+staff+" SET first_name=$1, last_name=$2, email=$3 WHERE staff_id = $4", "Resul", "Necefli", "necefliresul@gmail.com", 2)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update  oldu")
}

func GetResultsFromDB() {
	// result  table  da   ki   deyerlere elaqe
	rows, err := db.Query("SELECT first_name, last_name, email FROM " + staff)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var firstName string
		var lastName string
		var email string

		// rows.Scan ile  deyerleri  deyiskenlere  yoneltme
		err := rows.Scan(&firstName, &lastName, &email)
		if err != nil {
			log.Fatal(err)
		}

		// deyerleri  ekrana  yzdiraq
		fmt.Printf("First Name: %s, Last Name: %s, Email: %s\n", firstName, lastName, email)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// ana  funksiyamizda   diger  funksialarimza zeng edek
func main() {
	InitDB()
	defer db.Close()
	Insertgo.Into()
	Updatesql()
	Sqll.Deletesql()
	GetResultsFromDB()
}
