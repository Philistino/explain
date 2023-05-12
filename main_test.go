package main

import (
	_ "embed"
	"io"
	"os"
	"testing"
)

// https://stackoverflow.com/a/68978689

type Bingo struct {
}

func TestMainFn(t *testing.T) {
	os.Args = []string{"", ex1Cmd}
	urlBase = server.URL
	main()
	t.Error()
}

func TestMainNoArg(t *testing.T) {
	os.Args = []string{""}
	old := os.Stderr
	//create read and write pupe
	r, w, _ := os.Pipe()
	defer w.Close()
	defer r.Close()
	// set the stdout to the pipe
	os.Stderr = w
	// we excute the function
	main()
	// close the resource
	w.Close()
	// reset the stdout back to the orignal
	os.Stderr = old
	got, _ := io.ReadAll(r)

	want := noArgMsg
	if string(got) != want {
		t.Errorf("TestMainNoArg failed. want: %s, got: %s", want, got)
	}
}

func TestMainUsage(t *testing.T) {
	os.Args = []string{"-h"}
	old := os.Stderr
	r, w, _ := os.Pipe()
	defer w.Close()
	defer r.Close()
	os.Stderr = w

	Usage()

	w.Close()
	os.Stderr = old
	got, _ := io.ReadAll(r)

	want := helpMsg
	if string(got) != want {
		t.Errorf("TestMainNoArg failed. want: %s, got: %s", want, got)
	}
}

//go:embed test_fixtures/ex0.html
var ex0Html []byte

//go:embed test_fixtures/ex1.html
var ex1Html []byte

//go:embed test_fixtures/ex1Nested.html
var ex1NestedHtml []byte

var tt = []struct {
	name          string
	data          []byte
	wantCmdsLen   int
	wantExplsLen  int
	wantNestedLen int
	wantCmds      []CmdPart
	wantExpls     []Explanation
	wantNested    []string
}{
	{
		name:          "example_0",
		data:          ex0Html,
		wantCmdsLen:   8,
		wantExplsLen:  5,
		wantNestedLen: 0,
		wantCmds: []CmdPart{
			{
				HelpRef: "help-0",
				CmdPart: ":(){",
			},
			{
				HelpRef: "help-1",
				CmdPart: ":",
			},
			{
				HelpRef: "help-2",
				CmdPart: "|",
			},
			{
				HelpRef: "help-1",
				CmdPart: ":",
			},
			{
				HelpRef: "help-3",
				CmdPart: "&",
			},
			{
				HelpRef: "help-0",
				CmdPart: "}",
			},
			{
				HelpRef: "help-4",
				CmdPart: ";",
			},
			{
				HelpRef: "help-1",
				CmdPart: ":",
			},
		},
		wantExpls: []Explanation{
			{"help-0", "A shell function is an object that is called like a simple command and executes a compound command with a new set of positional parameters. Shell functions are declared as follows: name () compound-command [redirection] function name [()] compound-command [redirection] This defines a function named name. The reserved word function is optional. If the function reserved word is supplied, the parentheses are optional. The body of the function is the compound command compound-command (see Compound Commands above). That command is usually a list of commands between { and }, but may be any command listed under Compound Commands above. compound-command is executed whenever name is specified as the name of a simple command. Any redirections (see REDIRECTION below) specified when a function is defined are performed when the function is executed. The exit status of a function definition is zero unless a syntax error occurs or a readonly function with the same name already exists. When executed, the exit status of a function is the exit status of the last command executed in the body. (See FUNCTIONS below.)"},
			{"help-1", "call shell function ':'"},
			{"help-2", "Pipelines A pipeline is a sequence of one or more commands separated by one of the control operators | or |&amp;. The format for a pipeline is: [time [-p]] [ ! ] command [ [|⎪|&amp;] command2 ... ] The standard output of command is connected via a pipe to the standard input of command2. This connection is performed before any redirections specified by the command (see REDIRECTION below). If |&amp; is used, the standard error of command is connected to command2's standard input through the pipe; it is shorthand for 2>&1;&amp;1 |. This implicit redirection of the standard error is performed after any redirections specified by the command. The return status of a pipeline is the exit status of the last command, unless the pipefail option is enabled. If pipefail is enabled, the pipeline's return status is the value of the last (rightmost) command to exit with a non-zero status, or zero if all commands exit successfully. If the reserved word ! precedes a pipeline, the exit status of that pipeline is the logical negation of the exit status as described above. The shell waits for all commands in the pipeline to terminate before returning a value. If the time reserved word precedes a pipeline, the elapsed as well as user and system time consumed by its execution are reported when the pipeline terminates. The -p option changes the output format to that specified by POSIX. When the shell is in posix mode, it does not recognize time as a reserved word if the next token begins with a `-'. The TIMEFORMAT variable may be set to a format string that specifies how the timing information should be displayed; see the description of TIMEFORMAT under Shell Variables below. When the shell is in posix mode, time may be followed by a newline. In this case, the shell displays the total user and system time consumed by the shell and its children. The TIMEFORMAT variable may be used to specify the format of the time information. Each command in a pipeline is executed as a separate process (i.e., in a subshell)."},
			{"help-3", "If a command is terminated by the control operator &amp;, the shell executes the command in the background in a subshell. The shell does not wait for the command to finish, and the return status is 0."},
			{"help-4", "Commands separated by a ; are executed sequentially; the shell waits for each command to terminate in turn. The return status is the exit status of the last command executed."},
		},
		wantNested: []string{},
	},
	{
		name:          "example_1",
		data:          ex1Html,
		wantCmdsLen:   5,
		wantExplsLen:  3,
		wantNestedLen: 1,
		wantCmds: []CmdPart{
			{
				HelpRef: "help-0",
				CmdPart: "for user in $(cut -f1 -d: /etc/passwd); do"},
			{
				HelpRef: "help-2",
				CmdPart: "crontab(1)",
			},
			{
				HelpRef: "",
				CmdPart: "-u $user -l",
			},
			{
				HelpRef: "help-1",
				CmdPart: "2>/dev/null",
			},
			{
				HelpRef: "help-0",
				CmdPart: "; done",
			},
		},
		wantExpls: []Explanation{
			{"help-0", "for name [ [ in [ word ... ] ] ; ] do list ; done The list of words following in is expanded, generating a list of items. The variable name is set to each element of this list in turn, and list is executed each time. If the in word is omitted, the for command executes list once for each positional parameter that is set (see PARAMETERS below). The return status is the exit status of the last command that executes. If the expansion of the items following in results in an empty list, no commands are executed, and the return status is 0."},
			{"help-2", "maintain crontab files for individual users (Vixie Cron)"},
			{"help-1", "Before a command is executed, its input and output may be redirected using a special notation interpreted by the shell. Redirection may also be used to open and close files for the current shell execution environment. The following redirection operators may precede or appear anywhere within a simple command or may follow a command. Redirections are processed in the order they appear, from left to right. Redirecting Output Redirection of output causes the file whose name results from the expansion of word to be opened for writing on file descriptor n, or the standard output (file descriptor 1) if n is not specified. If the file does not exist it is created; if it does exist it is truncated to zero size. The general format for redirecting output is: [n]>gt;word If the redirection operator is >gt;, and the noclobber option to the set builtin has been enabled, the redirection will fail if the file whose name results from the expansion of word exists and is a regular file. If the redirection operator is >|t;|, or the redirection operator is >gt; and the noclobber option to the set builtin command is not enabled, the redirection is attempted even if the file named by word exists."},
		},
		wantNested: []string{"cut -f1 -d: /etc/passwd"},
	},
	{
		name:          "example_5",
		data:          ex1NestedHtml,
		wantCmdsLen:   4,
		wantExplsLen:  4,
		wantNestedLen: 0,
		wantCmds:      []CmdPart{{"help-0", "cut(1)"}, {"help-1", "-f1"}, {"help-2", "-d:"}, {"help-3", "/etc/passwd"}},
		wantExpls: []Explanation{
			{"help-0", "remove sections from each line of files"}, {"help-1", "-f, --fields=LIST select only these fields; also print any line that contains no delimiter character, unless the -s option is specified"}, {"help-2", "-d, --delimiter=DELIM use DELIM instead of TAB for field delimiter"}, {"help-3", "With no FILE, or when FILE is -, read standard input."},
		},
		wantNested: []string{},
	},
}
