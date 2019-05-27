package core

import "net/http"

type Message struct {
	StatusCode http.ConnState
	Message    string
}
