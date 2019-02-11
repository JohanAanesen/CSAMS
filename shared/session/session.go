package session

import "net/http"

type Session struct {
	Values map[string]interface{}
}

func Instance(r *http.Request) *Session {
	return &Session{}
}

