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

	want := "Shortest path to reach the vault (1): RLDUDRDDRR\n" +
		"Longest path length to reach the vault (2): 590\n"

	if string(out) != want {
		t.Errorf("\nWant:\n%s\nGot:\n%s", want, out)
	}
}
