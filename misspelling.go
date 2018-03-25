package spell

import (
	"strings"

	"github.com/chuckha/spell/parsers"
)

// This is a thing that might be out of order at some point so it needs to be fully expressed in this structure.
type Misspelling struct {
	// The misspelled word.
	Word string
	// The file the word is found in.
	File string
	// The line in the file the word is on.
	Line     int
	StartPos int // ??
}

// Start with just comments. Combines an AST and a dictionary to return a "misspelling"
func FindMisspellings(code *parsers.Code) ([]*Misspelling, error) {
	allWords := Default
	misspellings := []*Misspelling{}
	for _, comment := range code.Comments {
		for _, word := range strings.Fields(comment.Value) {
			if IsPunctuation(word) {
				allWords.Add(word)
			}

			if IsMisspelled(allWords, word) {
				misspellings = append(misspellings, &Misspelling{
					Word:     word,
					StartPos: comment.Pos,
				})
			}
		}
	}
	return misspellings, nil
}

func buildDict() WordSet {
	// get a custom dict of non words
	customTokens := map[string]struct{}{}
	return New(Default, customTokens)
}

// Update this for different dictionaries
func IsMisspelled(words WordSet, word string) bool {
	return !words.Exists(strings.ToLower(word))
}

// returns true if it's entirely punctuation
func IsPunctuation(word string) bool {
	return false
}
