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

	want := "Minimum number of steps with 5 pairs generator-microchip (1): 37\n" +
		"Minimum number of steps with 7 pairs generator-microchip (2): 61\n"

	if string(out) != want {
		t.Errorf("\nWant:\n%s\nGot:\n%s", want, out)
	}
}
