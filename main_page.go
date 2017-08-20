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
	list_file := []string{"", "ressources/header.html", "ressources/main_page.html", "ressources/footer.html"}
	list_bool := []bool{false, true, true, true}
	list_func := []fnc_page{stock_image, header, nil, footer}
	for i := range list_file {
		list_elem_page = append(list_elem_page, elem_page{list_file[i], list_bool[i], list_func[i]})
	}
	return list_elem_page
}

func stock_image(s string, user *session, r *http.Request) string {
	data := r.FormValue("data")
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

func main_page(w http.ResponseWriter, r *http.Request) {
	user := get_session(w, r)
	final, err := create_response(list_elem_main_page(), user, r)
	if err != nil {
		handle_err(err)
		return
	}
	w.Write([]byte(final))
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
