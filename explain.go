package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
)

func Explain(cmd string, baseUrl string) string {
	builder := strings.Builder{}
	var err error
	var f func(string)
	f = func(c string) {
		b, url, err := getPage(context.Background(), baseUrl, c)
		if err != nil {
			return
		}
		parsed := ParseReponse(b)
		if parsed.ErrorMsg != "" {
			fmt.Fprint(
				flag.CommandLine.Output(),
				"Error from explainshell.com: ",
				parsed.ErrorMsg,
			)
		}
		builder.WriteString(output(parsed.CmdParts, parsed.Expls, url))
		for _, nested := range parsed.NestedCmds {
			f(nested)
		}
	}
	f(cmd)

	if err != nil {
		return err.Error()
	}
	return builder.String()
}
