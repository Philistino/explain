package main

import (
	"strings"
	"testing"
)

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
