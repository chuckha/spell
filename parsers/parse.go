package parsers

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// Supported languages
	golang = "go"
	python = "python"
)

// Parse detects filetype and returns the parsed code that we care about.
func Parse(f *os.File) (*Code, error) {
	language, err := detectLanguage(f)
	if err != nil {
		return nil, fmt.Errorf("error detecting language: %v", err)
	}
	parserMap := map[string]Parser{
		golang: &GoParser{},
	}
	return parserMap[language].Parse(f)
}

// detectLanguage returns the best guess of the programming language contained in the file.
func detectLanguage(f *os.File) (string, error) {
	switch filepath.Ext(f.Name()) {
	case ".go":
		return golang, nil
	case ".py":
		return python, nil
	default:
		return "", fmt.Errorf("unknown language for file %v", f.Name())
	}
}
