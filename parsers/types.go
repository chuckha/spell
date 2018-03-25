package parsers

import (
	"os"
)

// Code is a representation of the code we care to spellcheck.
type Code struct {
	Path     string
	Comments []*Comment
}

// A comment from source
type Comment struct {
	// the position of the comment in the file
	Pos int
	// the value of the comment
	Value string
}

type Parser interface {
	Parse(*os.File) (*Code, error)
}
