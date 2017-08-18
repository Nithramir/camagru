package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func add_file(file string, w http.ResponseWriter) error {
	response, err := ioutil.ReadFile(file)
	if err != nil {
		handle_err(err)
		return err
	}
	fmt.Fprintln(w, string(response))
	return err
}

func index(w http.ResponseWriter, r *http.Request) {
	users, err := get_session(w, r)
	fmt.Printf("%v", users)
	err = add_file("ressources/header.html", w)
	if err != nil {
		w.Write([]byte("Erreur"))
		return
	}

	err = add_file("ressources/footer.html", w)
	if err != nil {
		w.Write([]byte("Erreur"))
		return
	}

	print("good")
}
