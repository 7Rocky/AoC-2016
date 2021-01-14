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

	want := "Number of pixels lit (1): 128\n" +
		"Message(2):\n\n" +
		"####..##...##..###...##..###..#..#.#...#.##...##..\n" +
		"#....#..#.#..#.#..#.#..#.#..#.#..#.#...##..#.#..#.\n" +
		"###..#..#.#..#.#..#.#....#..#.####..#.#.#..#.#..#.\n" +
		"#....#..#.####.###..#.##.###..#..#...#..####.#..#.\n" +
		"#....#..#.#..#.#.#..#..#.#....#..#...#..#..#.#..#.\n" +
		"####..##..#..#.#..#..###.#....#..#...#..#..#..##..\n"

	if string(out) != want {
		t.Errorf("\nWant:\n%s\nGot:\n%s", want, out)
	}
}
