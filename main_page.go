package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func list_elem_main_page() []elem_page {
	var list_elem_page []elem_page
	list_file := []string{"", "ressources/header.html", "ressources/main_page.html", "", "ressources/footer.html"}
	list_bool := []bool{false, true, true, false, true}
	list_func := []fnc_page{stock_image, header, content_main_page, show_all_images, footer}
	for i := range list_file {
		list_elem_page = append(list_elem_page, elem_page{list_file[i], list_bool[i], list_func[i]})
	}
	return list_elem_page
}

func content_main_page(s string, user *session, r *http.Request) string {
	if !user.info.is_connected {
		return "You need to connect to access this page"
	}
	final, _ := add_file("ressources/main_page.html")
	return final

}

func stock_image(s string, user *session, r *http.Request) string {
	if !user.info.is_connected {
		return ""
	}
	data := r.FormValue("data")
	if data != "" {
		data = strings.Replace(data, "data:image/png;base64,", "", 1)
		data2, err := b64.StdEncoding.DecodeString(string(data))
		if err != nil {
			handle_err(err)
			return ""
		}
		image_name := image_id() + ".png"
		insert_in_ddb(image_name, user)
		err = ioutil.WriteFile("ressources/img/"+image_name, data2, 0644)
		if err != nil {
			handle_err(err)
			return ""
		}
		return ""
	}
	return ""
}

func main_page(w http.ResponseWriter, r *http.Request) {
	user := get_session(w, r)
	final, err := create_response(list_elem_main_page(), user, r)
	if err != nil {
		handle_err(err)
		return
	}
	w.Write([]byte(final))
}

func show_all_images(s string, user *session, r *http.Request) string {
	var image_name string
	var final string

	if !user.info.is_connected {
		return ""
	}
	request := "SELECT name FROM " + table_image + ";"
	rows, err := datab.Query(request)
	if err != nil {
		handle_err(err)
		return "Erreur interne du serveur"
	}
	for rows.Next() {
		rows.Scan(&image_name)
		tmp := "<img src=\"" + "ressources/img/" + image_name + "\">"
		final += tmp

	}
	return final

}

func insert_in_ddb(image_name string, user *session) {
	request := "INSERT INTO " + table_image + " (name, author, aime, date) VALUE ('" + image_name + "', '" + user.info.pseudo + "', 0, NOW());"
	row, err := datab.Query(request)
	if err != nil {
		handle_err(err)
		return
	}
	row.Close()

}

func image_id() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return b64.URLEncoding.EncodeToString(b)
}
