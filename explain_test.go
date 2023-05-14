package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func TestExplainSuccess(t *testing.T) {
	t.Parallel()
	substrings := []string{
		`for name [ [ in [ word ... ] ] ; ] do list ; done The list of words following in is expanded, generating a list of items`,
		`maintain crontab files for individual users (Vixie Cron)`,
		"Unknown",
		"remove sections from each line of files",
		"With no FILE, or when FILE is -, read standard input.",
	}
	got := Explain(ex1Cmd, server.URL)
	for _, s := range substrings {
		if !strings.Contains(got, s) {
			t.Errorf("TestExplain failed. Missing: %s", s)
		}
	}
}
