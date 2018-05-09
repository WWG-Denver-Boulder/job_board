package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"net/http"

	"github.com/job_board/jbdb"
	"github.com/job_board/job_board_api/handlers/user_roles"
	"github.com/gorilla/mux"
)


// Set up main function with web server & router
func main() {	
	driver := "mysql"
	dataSource := "root:wwg@tcp(127.0.0.1:3306)/job_board_db"
	
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	jbDB := jbdb.New(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK") })
	
	urh := user_roles.New(jbDB)
	router.HandleFunc("/user_roles", urh.Get).Methods("GET")

	port := ":8080"
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
