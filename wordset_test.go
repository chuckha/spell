package spell_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/chuckha/spell"
)

func TestNewWords(t *testing.T) {
	testcases := []struct {
		name     string
		input    []byte
		expected map[string]struct{}
	}{
		{
			name:  "simple case",
			input: []byte("hello\nworld\nhow\nare\nyou"),
			expected: map[string]struct{}{
				"hello": struct{}{},
				"world": struct{}{},
				"how":   struct{}{},
				"are":   struct{}{},
				"you":   struct{}{},
			},
		},
		{
			name:     "empty case",
			input:    []byte(""),
			expected: map[string]struct{}{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := spell.NewWords(bytes.NewReader(tc.input))
			if err != nil {
				t.Fatalf("Did not expect an error here: %v", err)
			}
			for k := range actual {
				if _, ok := tc.expected[k]; !ok {
					fmt.Println(actual)
					fmt.Println(tc.expected)
					t.Fatalf("Expected to find %v but did not", k)
				}
			}
			if len(actual) != len(tc.expected) {
				t.Fatalf("Expected %v entries but got %v", len(tc.expected), len(actual))
			}
		})
	}
}

func TestNew(t *testing.T) {
	testCases := []struct {
		name     string
		wordSets []map[string]struct{}
		inputs   []string
		allFound bool
	}{
		{
			name: "easy case",
			wordSets: []map[string]struct{}{
				map[string]struct{}{
					"a": struct{}{},
				},
				map[string]struct{}{
					"b": struct{}{},
				},
			},
			inputs:   []string{"a", "b"},
			allFound: true,
		},
		{
			name: "all found fail case",
			wordSets: []map[string]struct{}{
				map[string]struct{}{
					"a": struct{}{},
				},
				map[string]struct{}{
					"b": struct{}{},
				},
			},
			inputs:   []string{"a", "c"},
			allFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ws := spell.New(tc.wordSets...)
			allFound := true
			for _, word := range tc.inputs {
				if !ws.Exists(word) {
					allFound = false
				}
			}
			if allFound != tc.allFound {
				t.Fatalf("Expected to find all of them? %v but we got %v", tc.allFound, allFound)
			}
		})

	}
}
