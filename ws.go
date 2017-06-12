package main

import (
	"database/sql"
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

func AgendaServer(w http.ResponseWriter, req *http.Request) {
	db, err = sql.Open("mysql", "root:admin@/agenda")
	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	rows, err := db.Query("SELECT nmAgenda FROM agenda")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	io.WriteString(w, "Agendas:\n")
	// Fetch rows
	for rows.Next() {
		var nmAgenda string
		rows.Scan(&nmAgenda)
		io.WriteString(w, nmAgenda+"\n")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

func main() {
	m := pat.New()
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))
	m.Get("/agenda/", http.HandlerFunc(AgendaServer))
	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	erro := http.ListenAndServe(":12345", nil)
	if erro != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

