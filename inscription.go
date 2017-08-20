package main

import (
	"net/http"
	"strings"
)

func list_elem_inscription() []elem_page {
	var list_elem_page []elem_page
	list_file := []string{"ressources/header.html", "ressources/inscription.html", "ressources/footer.html"}
	list_bool := []bool{true, true, true}
	list_func := []fnc_page{nil, nil, nil}
	for i := range list_file {
		list_elem_page = append(list_elem_page, elem_page{list_file[i], list_bool[i], list_func[i]})
	}
	return list_elem_page
}

func inscription(w http.ResponseWriter, r *http.Request) {
	user := get_session(w, r)
	//if user.info.is_connected == 1 {

	pseudo := r.FormValue("pseudo")
	password := r.FormValue("password")
	email := r.FormValue("email")
	if pseudo == "" || password == "" || email == "" {
		final, err := create_response(list_elem_inscription(), user, r)
		if err != nil {
			handle_err(err)
			return
		}
		w.Write([]byte(final))
	} else {
		data_inscription(pseudo, password, email)
		final, err := create_response(list_elem_inscription(), user, r)
		if err != nil {
			handle_err(err)
			return
		}
		final = strings.Replace(final, "PSEUDO", pseudo, -1)
		final = strings.Replace(final, "MAILDESTINATION", email, -1)
		w.Write([]byte(final))

	}
}
