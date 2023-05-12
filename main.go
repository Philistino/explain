package main

import (
	"flag"
	"fmt"
	"os"
)

const urlBase = "https://explainshell.com/explain"

const noArgMsg = `
Try passing an argument. For example: explain "chmod +x run.sh"

`

const helpMsg = `
explain is a tool to explain shell commands using information from explainshell.com right here in the terminal.
	
Usage:

	explain [command]

For example:

	explain "chmod +x run.sh"

`

var Usage = func() {
	fmt.Fprint(
		flag.CommandLine.Output(),
		helpMsg,
	)
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	cmd := flag.Arg(0)
	if cmd == "" {
		fmt.Fprint(
			flag.CommandLine.Output(),
			noArgMsg,
		)
		os.Exit(0)
	}

	fmt.Fprint(
		flag.CommandLine.Output(),
		Explain(cmd, urlBase),
	)
}
