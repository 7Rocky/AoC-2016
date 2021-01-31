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

	want := "Index producing the 64th one-time pad key (1): 15035\n" +
		"Index producing the 64th one-time pad key (2): 19968\n"

	if string(out) != want {
		t.Errorf("\nWant:\n%s\nGot:\n%s", want, out)
	}
}
