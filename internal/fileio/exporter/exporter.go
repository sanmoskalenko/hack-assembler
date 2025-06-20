package exporter

import (
	"os"
	"path/filepath"
	"strings"
)

type Exporter struct{}

func (r *Exporter) Export(file string, outDir string, data []byte) error {
	return Export(file, outDir, data)
}

func Export(fileName, destination string, content []byte) error {
	fn := toHackFileName(fileName)
	err := os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}
	path := filepath.Join(destination, fn)
	return os.WriteFile(path, content, 0644)
}

func toHackFileName(fileName string) string {
	if s := strings.Split(fileName, "."); len(s) == 2 {
		return s[0] + ".hack"
	}
	return fileName + ".hack"
}
