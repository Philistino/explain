package main

import (
	"bytes"
	"io"
	"log"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var whiteSpaceRgx = regexp.MustCompile(`\s+`)

var commandOrShellRgx = regexp.MustCompile(`command|shell`)

type CmdPart struct {
	HelpRef string
	CmdPart string
}

type Explanation struct {
	HelpRef string
	Help    string
}

func explanations(r io.Reader) []Explanation {
	tkn := html.NewTokenizer(r)
	expls := make([]Explanation, 0)
	var captureText bool
	var t html.Token
	var helpRef string
	whiteSpaceReplacement := []byte(" ")
	builder := strings.Builder{}
	for {
		switch tkn.Next() {

		case html.StartTagToken:
			t = tkn.Token()
			if t.Data != "pre" {
				continue
			}
			for _, a := range t.Attr {
				if a.Key == "id" {
					helpRef = a.Val
				}
			}
			captureText = true

		case html.TextToken:
			if !captureText {
				continue
			}
			t = tkn.Token()
			builder.Write(whiteSpaceRgx.ReplaceAll(tkn.Raw(), whiteSpaceReplacement))

		case html.EndTagToken:
			t = tkn.Token()
			if t.Data != "pre" {
				continue
			}
			expls = append(
				expls,
				Explanation{
					HelpRef: helpRef,
					Help:    builder.String(),
				},
			)
			builder.Reset()
			captureText = false

		// finished iteration
		case html.ErrorToken:
			return expls
		}
	}
}

func commands(r io.Reader) ([]CmdPart, []string) {
	n := findCommandDiv(r)
	if n == nil {
		return nil, nil
	}
	cmdParts := parseCommandDiv(n)
	nestedCmds := parseExpansionCmds(n)
	return cmdParts, nestedCmds
}

func ParseReponse(r []byte) ([]CmdPart, []Explanation, []string) {
	cmds, nestedCmds := commands(bytes.NewReader(r))
	exps := explanations(bytes.NewReader(r))
	return cmds, exps, nestedCmds
}

func findCommandDiv(r io.Reader) *html.Node {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	var returnNode *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "command" {
					returnNode = n
					break
				}
			}
		}
		if returnNode != nil {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return returnNode
}

// getHelpRef returns the helpref key if on this node.
func getHelpRef(node *html.Node) string {
	for _, a := range node.Attr {
		if a.Key == "helpref" {
			return a.Val
		}
	}
	return ""
}

func parseTextInNodeRecursively(n *html.Node) []string {
	var buf bytes.Buffer
	rw := io.ReadWriter(&buf)
	html.Render(rw, n)
	tkn := html.NewTokenizer(rw)
	strings := make([]string, 0)
	var t html.Token
	for {
		switch tkn.Next() {
		case html.TextToken:
			t = tkn.Token()
			strings = append(strings, whiteSpaceRgx.ReplaceAllString(t.Data, " "))

		// finished iteration
		case html.ErrorToken:
			return strings
		}
	}
}

func parseCommandDiv(node *html.Node) []CmdPart {
	var f func(*html.Node)
	cmds := make([]CmdPart, 0)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, a := range n.Attr {
				if a.Key == "class" && commandOrShellRgx.MatchString(a.Val) {
					cmds = append(cmds, CmdPart{
						CmdPart: strings.Join(parseTextInNodeRecursively(n), ""),
						HelpRef: getHelpRef(n),
					})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)
	return cmds
}

func parseExpansionCmds(node *html.Node) []string {
	var f func(*html.Node)
	cmds := make([]string, 0)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "title" && a.Val == "Zoom in to nested command" {
					cmds = append(cmds, strings.Join(parseTextInNodeRecursively(n), ""))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)
	return cmds
}
