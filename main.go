package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"simpleAPI/routs"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "karen"
	password = "1111"
	dbname   = "postgres"
)

func main() {

	router := mux.NewRouter()
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname =%s sslmode = disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStatement := `insert into "Users"("id", "firstname", "lastname", 
	 									"username", "email", "password_hushed")
	 									values($1, $2, $3, $4, $5, $6)
	 									RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/register", routs.UserRegister).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Request sample
// {
//     "firstname": "Karen",
//     "lastname": "Aghanyan",
//     "username": "Karenus",
//     "email": "karen@gmail.com",
//     "password": "112358",
//     "password_confirm": "112358"
// }
