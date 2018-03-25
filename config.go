package spell

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	spellConfigFilename = ".spellconfig"
)

var (
	DefaultConfig = &Config{}
)

type Config struct {
	// RootPath is the root of the search.
	RootPath string
	// Only return files that have an extension listed here.
	// If FileTypes is empty list all files.
	FileTypes []string
	// TODO improve this
	// IgnoreDirs are directories to ignore in the search. Will be used as a prefix match.
	IgnoreDirs []string
}

func LoadConfig(reader io.Reader) *Config {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		// TODO use logger
		fmt.Printf("error reding config reader: %v\n", err)
		return DefaultConfig
	}
	c := new(Config)
	err = json.Unmarshal(data, c)
	if err != nil {
		fmt.Printf("error unmarshalling config data: %v\n", err)
	}
	return c
}

type SearchPaths []string

func NewSearchPaths(paths ...string) *SearchPaths {
	s := make(SearchPaths, 0)
	for _, path := range paths {
		s.Add(path)
	}
	return &s
}

func (s *SearchPaths) Add(path string) {
	if path == "" {
		return
	}
	*s = append(*s, path)
}

// No merging of configs, just the first one
func FindConfig(path string) (string, error) {
	if path != "" {
		return path, nil
	}
	defaultPaths := NewSearchPaths(home(), curdir())
	spellConfigFile := ""
	for _, path := range *defaultPaths {
		// make sure it exists
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, spellConfigFilename) {
				spellConfigFile = path
			}
			return nil
		})
		if err != nil {
			return "", fmt.Errorf("error [%v] walking paths: %v", err, defaultPaths)
		}
		if spellConfigFile != "" {
			return spellConfigFile, nil
		}
	}
	return spellConfigFile, nil
}

func home() string {
	u, err := user.Current()
	if err != nil {
		fmt.Println("error getting homedir:", err)
		return ""
	}
	return filepath.Join(u.HomeDir, ".spell")
}

func curdir() string {
	d, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting cwd:", err)
		return ""
	}
	return d
}
