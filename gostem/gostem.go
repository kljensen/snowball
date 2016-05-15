//
// Creates a binary `gostem` that stems an input file.
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/kljensen/snowball"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	var language *string = flag.String("l", "english", "Language")
	var infile *string = flag.String("i", "", "Input file for stemming")
	flag.Parse()

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}

	bf := bufio.NewReader(f)

	for {
		line, isPrefix, err := bf.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", f.Name())
		}

		word := strings.TrimSpace(string(line))
		stemmed, err := snowball.Stem(word, *language, true)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(stemmed)
	}
}
