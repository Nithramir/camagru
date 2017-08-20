package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func handle_err(err error) {
	s := err.Error()
	fmt.Printf(
		"version: %s\ntypes: %T / %T\nstring value via err.Error(): %q\n",
		runtime.Version(), err, s, s)
}

func main() {

	init_database("root", "1234")
	http.HandleFunc("/index.html", index)
	http.HandleFunc("/database.php", database)
	http.HandleFunc("/inscription.html", inscription)
	http.HandleFunc("/deconnection.html", deconnection)
	http.HandleFunc("/connection.html", connection)
	http.HandleFunc("/main_page.html", main_page)

	http.Handle("/", http.StripPrefix("/ressources/", http.FileServer(http.Dir("./ressources/"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
