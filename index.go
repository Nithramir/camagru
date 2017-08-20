package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	users := get_session(w, r)
	fmt.Printf("%v", users)
	response, err := add_file("ressources/header.html")
	if err != nil {
		w.Write([]byte("Erreur"))
		return
	}

	response2, err := add_file("ressources/footer.html")
	if err != nil {
		w.Write([]byte("Erreur"))
		return
	}
	response += response2
	fmt.Fprintf(w, response)

	print("good")
}
