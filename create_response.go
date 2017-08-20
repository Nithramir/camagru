package main

import (
	"io/ioutil"
	"net/http"
)

type fnc_page func(s string, user *session, r *http.Request) string

type elem_page struct {
	list    string
	is_file bool
	ptrfunc fnc_page
}

func create_response(f []elem_page, user *session, r *http.Request) (string, error) {
	var response string
	var responsetmp string
	var err error
	for i := range f {
		if f[i].is_file {
			responsetmp, err = add_file(f[i].list)
			if err != nil {
				handle_err(err)
				return "", err
			}
		} else {
			responsetmp = f[i].list
		}
		if f[i].ptrfunc != nil {
			responsetmp = f[i].ptrfunc(responsetmp, user, r)
		}

		response += responsetmp
	}
	return response, nil
}

func add_file(file string) (string, error) {
	response, err := ioutil.ReadFile(file)
	if err != nil {
		handle_err(err)
		return "", err
	}
	return string(response), err
}
