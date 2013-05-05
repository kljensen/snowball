package main

import (
	"bufio"
	"github.com/kljensen/snowball/english"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		input, err := in.ReadString('\n')

		if err == io.EOF {
			break
		}
		_, err = out.WriteString(english.Stem(input) + "\n")
		if err != nil {
			panic(err)
		}
	}
}
