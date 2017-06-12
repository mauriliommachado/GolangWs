package main

import (
	"database/sql"
	"fmt"
	"github.com/bmizerany/pat"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func main() {
	iniciaBD()
	m := pat.New()
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	erro := http.ListenAndServe(":12345", nil)
	if erro != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func iniciaBD() {
	db, err = sql.Open("mysql", "root:admin@/agenda")
	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	var nmAgenda string
	err := db.QueryRow("SELECT nmAgenda FROM agenda").Scan(&nmAgenda)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(nmAgenda)
}
