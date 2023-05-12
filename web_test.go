package main

import (
	"net/http"
	"net/http/httptest"
)

const (
	ex0Cmd       = ":(){ :|:& };:"
	ex1Cmd       = "for user in $(cut -f1 -d: /etc/passwd); do crontab -u $user -l 2>/dev/null; done"
	ex1NestedCmd = "cut -f1 -d: /etc/passwd"
)

var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	cmd := r.URL.Query().Get("cmd")
	switch {
	case cmd == ex0Cmd:
		w.Write(ex0Html)
	case cmd == ex1Cmd:
		w.Write(ex1Html)
	case cmd == ex1NestedCmd:
		w.Write(ex1NestedHtml)
	}
}))
