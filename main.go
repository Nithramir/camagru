package main

import (
	"fmt"
	"net/http"
	"runtime"
	"log"
)

func handle_err(err error) {
	s := err.Error()
	fmt.Printf(
		"version: %s\ntypes: %T / %T\nstring value via err.Error(): %q\n",
		runtime.Version(), err, s, s)
}

func main() {

	http.HandleFunc("/index.html", index)
	http.HandleFunc("/database.php", database)
	http.Handle("/", http.StripPrefix("/ressources/img", http.FileServer(http.Dir("./ressources/img"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
