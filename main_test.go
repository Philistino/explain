package main

import (
	_ "embed"
	"io"
	"os"
	"testing"
)

//go:embed test_fixtures/ex0.html
var ex0Html []byte

//go:embed test_fixtures/ex1.html
var ex1Html []byte

//go:embed test_fixtures/ex1Nested.html
var ex1NestedHtml []byte

// func TestMainFn(t *testing.T) {
// 	os.Args = []string{"", ex1Cmd}
// 	urlBase = server.URL
// 	main()
// 	t.Error()
// }

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
