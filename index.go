package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func add_file(file string) (string, error) {
	response, err := ioutil.ReadFile(file)
	return string(response), err
}

func index(w http.ResponseWriter, r *http.Request) {
	users, err := get_session(w, r)
	fmt.Printf("%v", users)
	response, err := add_file("ressources/header.html")
	if err != nil {
		handle_err(err)
		return
	}

	fmt.Fprintln(w, string(response))
	response, err = add_file("ressources/footer.html")
	if err != nil {
		handle_err(err)
		return
	}

	fmt.Fprintln(w, string(response))

	print("good")
}
