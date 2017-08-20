package main

import "net/http"

func footer(s string, user *session, r *http.Request) string {
	var final string
	var err error
	final, err = add_file("ressources/footer.html")
	if err != nil {
		handle_err(err)
		return ""
	}
	return final

}
