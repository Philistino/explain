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

func TestParse(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name      string
		data      []byte
		wantCmds  int
		wantExpls int
	}{
		{
			name:      "example_0",
			data:      example_0,
			wantCmds:  8,
			wantExpls: 5,
		},
		{
			name:      "example_1",
			data:      example_1,
			wantCmds:  5,
			wantExpls: 3,
		},
		{
			name:      "example_2",
			data:      example_2,
			wantCmds:  1,
			wantExpls: 1,
		},
		{
			name:      "example_3",
			data:      example_3,
			wantCmds:  13,
			wantExpls: 6,
		},
		{
			name:      "example_4",
			data:      example_4,
			wantCmds:  10,
			wantExpls: 9,
		},
	}
	for _, tc := range tt {
		cmds, expls := ParseReponse(tc.data)
		if len(cmds) != tc.wantCmds {
			t.Errorf("TestParse failed on number of cmds: name %s, wanted %d, got %d", tc.name, tc.wantCmds, len(cmds))
		}
		if len(expls) != tc.wantExpls {
			t.Errorf("TestParse failed on number of cmds: name %s, wanted %d, got %d", tc.name, tc.wantExpls, len(expls))
		}
	}
}

func TestParseExplanations(t *testing.T) {
	cmds := Explanations(strings.NewReader(testHtml))
	if len(cmds) != 2 {
		t.Error()
	}
	for _, s := range cmds {
		log.Println(s)
	}
}

func TestFindCommandDiv(t *testing.T) {
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
	cmds := Commands(strings.NewReader(testHtml))
	if len(cmds) != 2 {
		t.Error()
	}
	for _, s := range cmds {
		log.Println(s)
	}
}

func BenchmarkExplanations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Explanations(strings.NewReader(testHtml))
	}
}
