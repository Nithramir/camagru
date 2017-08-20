package main

import "net/http"

func list_elem_deconnection() []elem_page {
	var list_elem_page []elem_page
	list_file := []string{"ressources/header.html", "Deconnection successfull", "ressources/footer.html"}
	list_bool := []bool{true, false, true}
	list_func := []fnc_page{header, nil, footer}
	for i := range list_file {
		list_elem_page = append(list_elem_page, elem_page{list_file[i], list_bool[i], list_func[i]})
	}
	return list_elem_page
}

func deconnection(w http.ResponseWriter, r *http.Request) {
	user := get_session(w, r)
	user.info.is_connected = false
	user.info.pseudo = ""
	final, _ := create_response(list_elem_deconnection(), user, r)
	w.Write([]byte(final))
	//http.Redirect(w, r, "http://localhost:8080/index.html", 301)
}
