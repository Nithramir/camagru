package main

import "net/http"

func inscription_file(w http.ResponseWriter) error {
	err := add_file("ressources/header.html", w)
	if err != nil {
		w.Write([]byte("Erreur"))
		handle_err(err)
	}
	err = add_file("ressources/inscription.html", w)
	if err != nil {
		w.Write([]byte("Erreur"))
		handle_err(err)
	}
	err = add_file("ressources/footer.html", w)
	if err != nil {
		w.Write([]byte("Erreur"))
		handle_err(err)
	}
	return err
}

func inscription(w http.ResponseWriter, r *http.Request) {
	pseudo := r.FormValue("peudo")
	password := r.FormValue("password")
	email := r.FormValue("email")
	if pseudo == "" || password == "" || email == "" {
		err := inscription_file(w)
		if err != nil {
			return
		}
	}
}
