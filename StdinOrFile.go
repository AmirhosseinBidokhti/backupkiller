package main

import (
	"bufio"
	"io"
	"os"
)

func StdinOrFile(readChoice string) []string {
	var in io.Reader

	if readChoice == "stdin" {
		in = os.Stdin

	} else {
		filename := readChoice
		f, err := os.Open(filename)

		if err != nil {
			panic(err)
		}
		defer f.Close()
		in = f
	}

	sc := bufio.NewScanner(in)

	if err := sc.Err(); err != nil {
		panic(err)
	}

	var data []string
	for sc.Scan() {
		data = append(data, sc.Text())
	}

	return data

}
