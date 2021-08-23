package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"random-text-generator/generator"
	"random-text-generator/trigram"
)

const usage = `
Usage: %s [corpus.txt] [n]

This program parses all words from a given text corpus to randomly generate a
new text that is looking similar to the provided text.
`

type Solution interface {
	Add(trigram.Trigram)
	Generate(n int) (string, error)
}

func main() {
	// Initialize the RNG so we get a different output on each run.
	rand.Seed(time.Now().UnixNano())

	corpusFile, n := parseFlags()
	trigrams, err := trigram.ParseFile(corpusFile)
	if err != nil {
		log.Fatal(err)
	}

	var s Solution = generator.New()

	// We pass all parsed trigrams to your solution so it can learn its model
	// based on the provided corpus of words.
	for _, t := range trigrams {
		s.Add(t)
	}

	// We now want to generate a new text based on the trigrams that we added above.
	result, err := s.Generate(n)
	if err != nil {
		log.Fatal(err)
	}

	// Finally we print the result to stdout.
	fmt.Println(result)
}

func parseFlags() (corpusFile string, n int) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	corpusFile = "corpus/trump.txt"
	if len(args) > 0 {
		corpusFile = args[0]
	}

	n = 100
	if len(args) > 1 {
		var err error
		n, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("second argument is not an integer: %v", err)
		}
	}

	return corpusFile, n
}
