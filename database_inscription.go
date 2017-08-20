package main

func data_inscription(pseudo string, password string, email string) error {
	request := "INSERT INTO " + table_name + "(pseudo, password, email, date) VALUE ('" + pseudo + "', MD5('" + password + "'), '" + email + "', NOW());"
	print(request)
	req, err := datab.Query(request)
	if err != nil {
		handle_err(err)
		return err
	}
	req.Close()
	return nil
}

func database_connection(pseudo string, password string) bool {
	request := "SELECT pseudo FROM " + table_name + " WHERE pseudo = '" + pseudo + "' AND password = MD5('" + password + "');"
	print(request)
	req, err := datab.Query(request)
	if err != nil {
		handle_err(err)
		return false
	}
	defer req.Close()
	for req.Next() {
		return true
	}
	return false

}
