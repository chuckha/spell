package parsers

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

type GoParser struct{}

// Parse parses the file passed in assuming it's a go file and returns a code representation.
func (g *GoParser) Parse(f *os.File) (*Code, error) {
	fset := token.NewFileSet()
	tree, err := parser.ParseFile(fset, f.Name(), f, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %v", err)
	}
	code := &Code{
		Path:     f.Name(),
		Comments: []*Comment{},
	}

	for _, comment := range tree.Comments {
		for _, cmt := range comment.List {
			code.Comments = append(code.Comments, &Comment{
				Pos:   int(cmt.Pos()),
				Value: cmt.Text,
			})
		}
	}
	return code, nil
}
