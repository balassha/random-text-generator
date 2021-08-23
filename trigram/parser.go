package trigram

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// ParseFile parses all trigrams from a given file.
func ParseFile(path string) ([]Trigram, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return Parse(f)
}

// Parse reads trigrams from an io.Reader.
func Parse(r io.Reader) ([]Trigram, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(words) < 3 {
		return nil, errors.New("input has less than 3 words")
	}

	var trigrams []Trigram
	for i := 0; i < len(words)-2; i++ {
		trigram := Trigram{words[i], words[i+1], words[i+2]}
		trigrams = append(trigrams, trigram)
	}

	return trigrams, nil
}
