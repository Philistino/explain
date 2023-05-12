package main

import (
	"context"
	"strings"
)

func Explain(cmd string, baseUrl string) string {
	builder := strings.Builder{}
	var err error
	var f func(string)
	f = func(c string) {
		b, url, err := GetPage(context.Background(), baseUrl, c)
		if err != nil {
			return
		}
		cmds, expls, nestedCmds := ParseReponse(b)
		builder.WriteString(Output(cmds, expls, url))
		for _, nested := range nestedCmds {
			f(nested)
		}
	}
	f(cmd)

	if err != nil {
		return err.Error()
	}
	return builder.String()
}
