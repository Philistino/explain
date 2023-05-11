package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

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
	urlBase := "https://explainshell.com/explain"
	b, err := GetPage(context.Background(), urlBase, cmd)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	cmds, expls := ParseReponse(b)
	fmt.Print(Output(cmds, expls))
}
