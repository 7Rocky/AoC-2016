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

	want := "Bathroom code (1): 78985\n" +
		"Bathroom code (2): 57DD8\n"

	if string(out) != want {
		t.Errorf("Want:\n%s\nGot:\n%s", want, out)
	}
}
