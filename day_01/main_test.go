package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	want := "Blocks (1): 287\n" +
		"Blocks to twice reached position (2): 133\n"

	if string(out) != want {
		t.Errorf("Want:\n%s\nGot:\n%s", want, out)
	}
}
