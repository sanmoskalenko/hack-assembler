package loader

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name string
	Data []string
}

type Loader struct{}

func (r *Loader) LoadASM(dir string) ([]File, error) {
	return LoadASM(dir)
}

const asmExt = ".asm"

var ErrUnsupportedExt = errors.New("unsupported file extension")

func LoadASM(path string) ([]File, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return loadFromDir(path)
	}
	if filepath.Ext(path) != asmExt {
		return nil, ErrUnsupportedExt
	}

	data, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return []File{{Name: filepath.Base(path), Data: data}}, nil
}

func loadFromDir(path string) ([]File, error) {
	files, err := collectAssemblerFiles(path)
	if err != nil {
		return nil, err
	}
	var result []File
	for _, fn := range files {
		d, err := readFile(filepath.Join(path, fn))
		if err != nil {
			return nil, err
		}
		result = append(result, File{Name: fn, Data: d})
	}
	return result, nil
}

func readFile(path string) ([]string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(f), "\n"), nil
}

func collectAssemblerFiles(dn string) ([]string, error) {
	entries, err := os.ReadDir(dn)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fileName := entry.Name()
		if filepath.Ext(fileName) == asmExt {
			files = append(files, fileName)
		}
	}
	return files, nil
}
