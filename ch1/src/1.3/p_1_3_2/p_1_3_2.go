package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filename := ""
	if len(files) == 0 {
		filename = "standard input"
		countLines(os.Stdin, counts, filename)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error while Opening %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

}

func countLines(f *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filename)
		}
	}
}
