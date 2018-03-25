package spell

import (
	"os"
	"path/filepath"
	"strings"
)

type FileFinder struct {
	Config *Config
}

func NewFileFinder(cfg *Config) *FileFinder {
	return &FileFinder{
		Config: cfg,
	}
}

// Find returns a list of files based on the configuration
func (f *FileFinder) Find() ([]string, error) {
	filesToCheck := make([]string, 0)
	err := filepath.Walk(f.Config.RootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, suffix := range f.Config.FileTypes {
			if strings.HasSuffix(path, suffix) {
				filesToCheck = append(filesToCheck, path)
			}
		}
		return nil
	})
	return filesToCheck, err
}
