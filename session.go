package main

import ("io"
"crypto/rand"
"encoding/base64"
"net/http"
)

var list_session []*session

const(session_name = "go_session")

type information struct {
	is_connected int
}

type session struct {	
	cookieName  string     //private cookiename
	cookie_id   string
	maxlifetime int64
	info 		information
}

func sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func create_new_session() *session {
	new_session :=  session{session_name, sessionId(), 150, information{0} }
	list_session = append(list_session, &new_session)
	return &new_session
}

func find_session(cookie string) *session{
	for i := range list_session {
		if list_session[i].cookie_id == cookie {
			return list_session[i]
		}
	}
	return nil
}

func get_session(w http.ResponseWriter, r *http.Request) (*session, error){
	cookie_session, err := r.Cookie(session_name)
	if err != nil {
		create_sess := create_new_session()	
		cookie := http.Cookie{Name: session_name, Value: create_sess.cookie_id, Path: "/", HttpOnly: true, MaxAge: int(create_sess.maxlifetime)}
		http.SetCookie(w, &cookie)
		return create_sess, nil
	} else {
		find_sess := find_session(cookie_session.Value)
		if find_sess != nil {
			return find_sess, nil
		} else {

			return create_new_session(), nil	
		}

	}
}
