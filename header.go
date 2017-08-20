package main

import (
	"net/http"
	"strings"
)

func header(s string, user *session, r *http.Request) string {
	var final string
	var connected string
	var err error
	final, err = add_file("ressources/header.html")
	if err != nil {
		handle_err(err)
		return ""
	}
	if user.info.is_connected {
		connected, err = add_file("ressources/is_connected.html")
		if err != nil {
			handle_err(err)
			return ""
		}
		connected = strings.Replace(connected, "PSEUDO", user.info.pseudo, 1)
		return final + connected
	} else {
		connected, err = add_file("ressources/not_connected.html")
		if err != nil {
			handle_err(err)
			return ""
		}
		connected = strings.Replace(connected, "pseudo", user.info.pseudo, 1)
		return final + connected

	}
}
