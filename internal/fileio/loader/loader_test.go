package loader_test

import (
	"errors"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/sanmoskalenko/hack-assembler/internal/fileio/loader"
)

const testDir = "testdata"

func TestLoadSingleValidAsmFile(t *testing.T) {
	path := filepath.Join(testDir, "valid.asm")
	files, err := loader.LoadASM(path)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}
	if files[0].Name != "valid.asm" {
		t.Fatalf("expected file name `valid.asm`, got %s", files[0].Name)
	}

	want := []string{"@2", "D+1"}
	got := files[0].Data
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected slice: got %v, want %v", got, want)
	}
}

func TestLoadSingleInvalidFile(t *testing.T) {
	path := filepath.Join(testDir, "ignore.txt")
	_, err := loader.LoadASM(path)
	if !errors.Is(err, loader.ErrUnsupportedExt) {
		t.Fatalf("expected ErrUnsupportedExt, got %v", err)
	}
}

func TestLoadFromDirectory(t *testing.T) {
	files, err := loader.LoadASM(testDir)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(files) != 2 {
		t.Errorf("expected 2 asm files, got %d", len(files))
	}

	got := files
	want := []loader.File{
		{Name: "valid.asm", Data: []string{"@2", "D+1"}},
		{Name: "empty.asm", Data: []string{""}},
	}
	sort.Slice(got, func(i, j int) bool { return got[i].Name < got[j].Name })
	sort.Slice(want, func(i, j int) bool { return want[i].Name < want[j].Name })
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected slice: got %v, want %v", got, want)
	}
}
