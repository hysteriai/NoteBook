//Exercise 1.4: Modify dup2 to print the names of all files in which each dup licated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countsFile := make(map[string]int)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, arg, countsFile)
		f.Close()
	}

	for file, n := range countsFile {
		if n > 1 {
			fmt.Printf("Dup line file: %s\n", file)
		}
	}
}

func countLines(f *os.File, filename string, countsFile map[string]int) {
	countsLine := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		countsLine[input.Text()]++
	}
	for _, n := range countsLine {
		if n > 1 {
			countsFile[filename]++
		}
	}
}
