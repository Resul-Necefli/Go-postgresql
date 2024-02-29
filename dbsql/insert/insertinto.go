package Insertgo

///5555
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
	password = "*********"
	dbname   = "DVDrental"
	result   = "result"
)

var DB *sql.DB

func Into() {
	Cenvres := fmt.Sprintf("host=%s port=%s user=%s password =%s dbname =%s  sslmode = disable", host, port, user, password, dbname)

	DB, err := sql.Open("postgres", Cenvres)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	fmt.Println("kayt islemi  basladi")
	var name string = "Resul"
	var last_name string = "Necefli"
	var email string = "necefliresul29@gmail.com"
	_, err = DB.Exec("insert into result(first_name,last_name,email) values ($1, $2 ,$3)", name, last_name, email)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert islemi  tamamlandi")
	}

}
