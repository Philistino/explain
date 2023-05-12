package main

import (
	_ "embed"
	"log"
	"strings"
	"testing"
)

//go:embed explainshell.htm
var testHtml string

//go:embed explainshell_long.htm
var testHtmlLong string

//go:embed fixtures/example_0.html
var example_0 []byte

//go:embed fixtures/example_1.html
var example_1 []byte

//go:embed fixtures/example_2.html
var example_2 []byte

//go:embed fixtures/example_3.html
var example_3 []byte

//go:embed fixtures/example_4.html
var example_4 []byte

//go:embed fixtures/example_5.html
var example_5 []byte

func TestParse(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name       string
		data       []byte
		wantCmds   int
		wantExpls  int
		wantNested int
	}{
		{
			name:       "example_0",
			data:       example_0,
			wantCmds:   8,
			wantExpls:  5,
			wantNested: 0,
		},
		{
			name:       "example_1",
			data:       example_1,
			wantCmds:   5,
			wantExpls:  3,
			wantNested: 1,
		},
		{
			name:       "example_2",
			data:       example_2,
			wantCmds:   1,
			wantExpls:  1,
			wantNested: 1,
		},
		{
			name:       "example_3",
			data:       example_3,
			wantCmds:   13,
			wantExpls:  6,
			wantNested: 0,
		},
		{
			name:       "example_4",
			data:       example_4,
			wantCmds:   10,
			wantExpls:  9,
			wantNested: 0,
		},
	}
	for _, tc := range tt {
		cmds, expls, nested := ParseReponse(tc.data)
		if len(cmds) != tc.wantCmds {
			t.Errorf("TestParse failed on number of cmds: name %s, wanted %d, got %d", tc.name, tc.wantCmds, len(cmds))
		}
		if len(expls) != tc.wantExpls {
			t.Errorf("TestParse failed on number of cmds: name %s, wanted %d, got %d", tc.name, tc.wantExpls, len(expls))
		}
		if len(nested) != tc.wantNested {
			t.Errorf("TestParse failed on number of nested cmds: name %s, wanted %d, got %d", tc.name, tc.wantNested, len(nested))
		}
	}
}

func TestParseExplanations(t *testing.T) {
	t.Parallel()
	cmds := explanations(strings.NewReader(testHtml))
	if len(cmds) != 2 {
		t.Error()
	}
	for _, s := range cmds {
		log.Println(s)
	}
}

func TestFindCommandDiv(t *testing.T) {
	t.Parallel()
	cmd := findCommandDiv(strings.NewReader(testHtmlLong))
	if cmd.Data != "div" {
		t.Error("TestFind did not find the correct div", cmd.Data)
	}
	commandAttrExists := false
	for _, a := range cmd.Attr {
		if a.Val == "command" {
			commandAttrExists = true
			break
		}
	}
	if !commandAttrExists {
		t.Error("TestFind did not find the div with the correct attr")
	}
}

func TestCommand(t *testing.T) {
	t.Parallel()
	cmds, nested := commands(strings.NewReader(testHtml))
	if len(cmds) != 2 {
		t.Error()
	}
	if len(nested) != 0 {
		t.Error()
	}
	for _, s := range cmds {
		log.Println(s)
	}
}
