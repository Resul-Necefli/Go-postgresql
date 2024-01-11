package Sqll

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "DVDrental "
	password = "yenisifre"
	result   = "result"
)

func Deletesql() {
	datavalues := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode = disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", datavalues)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("baglanti qurldu")
	id := 1
	_, err = db.Exec("DELETE FROM "+result+" WHERE ID = $1", id)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("silme  islemi tamamlandi")

}
