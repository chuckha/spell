package spell_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/chuckha/spell"
)

func TestFind(t *testing.T) {
	files := []string{"main.go", "test.py", "one.sh"}
	expected := []string{"main.go", "test.py"}
	WithTempDir(t, func(tmp string) {
		fullExpectedPaths := make([]string, len(expected))
		for i, exp := range expected {
			fullExpectedPaths[i] = filepath.Join(tmp, exp)
		}
		for _, file := range files {
			ioutil.WriteFile(filepath.Join(tmp, file), []byte{}, 0)
		}

		f := spell.FileFinder{
			Config: &spell.Config{
				RootPath: tmp,
				FileTypes: []string{
					".go",
					".py",
				},
			},
		}

		files, err := f.Find()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(fullExpectedPaths, files) {
			t.Fatalf("expected %v got %v", expected, files)
		}
	})
}

func WithTempDir(t *testing.T, fn func(tmpdir string)) {
	tmp := os.TempDir()
	out, err := ioutil.TempDir(tmp, "spell-test")
	if err != nil {
		t.Fatal(err)
	}
	fn(out)
	defer os.RemoveAll(out)
}
