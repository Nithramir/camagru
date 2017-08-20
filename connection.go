package main

import (
	"net/http"
	"strings"
)

func list_elem_connection() []elem_page {
	var list_elem_page []elem_page
	list_file := []string{"", "ressources/header.html", "ressources/connection.html", "ressources/footer.html"}
	list_bool := []bool{false, true, true, true}
	list_func := []fnc_page{connection_success, header, form_connection, footer}
	for i := range list_file {
		list_elem_page = append(list_elem_page, elem_page{list_file[i], list_bool[i], list_func[i]})
	}
	return list_elem_page
}

func form_connection(s string, user *session, r *http.Request) string {
	var final string
	pseudo := r.FormValue("pseudo")
	password := r.FormValue("password")
	if user.info.is_connected {
		return "utilisateur déjà connecté, vous pouvez vous déconnecter"
	}
	if pseudo == "" || password == "" {
		final, _ = add_file("ressources/connection.html")
		return final
	}
	if database_connection(pseudo, password) == true {
		user.info.is_connected = true
		user.info.pseudo = pseudo
		final, _ = add_file("ressources/connection_successfull.html")
		final = strings.Replace(final, "PSEUDO", pseudo, 1)
		print("connection successfull")
		return final

	} else {
		return ""
	}
}

func connection_success(s string, user *session, r *http.Request) string {
	pseudo := r.FormValue("pseudo")
	password := r.FormValue("password")
	if user.info.is_connected {
		return ""
	}
	if pseudo == "" || password == "" {
		return ""
	}
	if database_connection(pseudo, password) == true {
		user.info.is_connected = true
		user.info.pseudo = pseudo
		return ""
	}
	return ""
}

func connection(w http.ResponseWriter, r *http.Request) {
	user := get_session(w, r)

	final, _ := create_response(list_elem_connection(), user, r)
	w.Write([]byte(final))

}
