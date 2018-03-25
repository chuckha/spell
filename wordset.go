package spell

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	dictFile = "/usr/share/dict/words"
)

var (
	Default WordSet
)

func init() {
	words, err := UsrShareDictWords()
	if err != nil {
		panic(err)
	}
	Default = New(words)
}

type WordSet map[string]struct{}

func (w WordSet) Exists(word string) bool {
	_, ok := w[word]
	return ok
}

func (w WordSet) Add(word string) {
	w[word] = struct{}{}
}

func New(wordSets ...map[string]struct{}) WordSet {
	words := make(map[string]struct{})
	// Merge the wordSets brought in
	for _, wordSet := range wordSets {
		for word := range wordSet {
			w := strings.ToLower(word)
			words[w] = struct{}{}
		}
	}
	return words
}

func UsrShareDictWords() (map[string]struct{}, error) {
	f, err := os.Open(dictFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewWords(f)
}

func NewWords(reader io.Reader) (map[string]struct{}, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	dictionary := make(map[string]struct{})
	for _, wordBytes := range bytes.Split(bytes.TrimSpace(b), []byte("\n")) {
		if len(wordBytes) == 0 {
			continue
		}
		dictionary[string(wordBytes)] = struct{}{}
	}
	return dictionary, nil
}
