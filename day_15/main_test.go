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

	want := "Time to press the button (1): 122318\n" +
		"Time to press the button with an extra disc (2): 3208583\n"

	if string(out) != want {
		t.Errorf("\nWant:\n%s\nGot:\n%s", want, out)
	}
}
