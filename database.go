package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"runtime"
)

var datab *sql.DB

const (
	table_name = "camagru"
	ddb_name   = "table1"
)

func database(w http.ResponseWriter, r *http.Request) {
	_, err := init_database("root", "1234")
	if err != nil {
		s := err.Error()
		fmt.Fprintf(w,
			"version: %s\ntypes: %T / %T\nstring value via err.Error(): %q\n",
			runtime.Version(), err, s, s)

		return
	}

}

func init_database(db_user string, db_passw string) (*sql.DB, error) {
	db, err := sql.Open("mysql", db_user+":"+db_passw+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		handle_err(err)
		return nil, err
	}
	_, err = db.Exec("CREATE DATABASE if not exists " + ddb_name)
	if err != nil {
		handle_err(err)
		return nil, err
	}

	_, err = db.Exec("USE " + ddb_name)
	if err != nil {
		handle_err(err)
		return nil, err
	}

	fmt.Println("Connected to database " + ddb_name)
	rows, err := db.Query("CREATE TABLE if not exists " + table_name + "(ID int NOT NULL AUTO_INCREMENT, pseudo CHAR(255), password CHAR(255), email char (255), date DATETIME);")
	if err != nil {
		handle_err(err)
		return nil, err
	}

	fmt.Println("Connected to " + table_name)
	defer rows.Close()
	return db, err
}