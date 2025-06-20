package symboltable_test

import (
	"strconv"
	"testing"

	"github.com/sanmoskalenko/hack-assembler/internal/symboltable"
)

func TestPredefinedSymbols(t *testing.T) {
	sut := symboltable.New()
	tests := map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"SCREEN": 16384,
		"KBD":    24576,
	}
	for i := range 16 {
		tests["R"+strconv.Itoa(i)] = i
	}

	for s, exp := range tests {
		if !sut.Contains(s) {
			t.Errorf("expected symbol table to contain %s", s)
			continue
		}

		addr := sut.GetAddress(s)
		if addr != exp {
			t.Errorf("for symbol %s, expected address %d, got %d", s, exp, addr)
		}
	}
}

func TestAddAndGetVar(t *testing.T) {
	sut := symboltable.New()

	addr := sut.AddVarEntry("foo")
	if addr != 16 {
		t.Errorf("expected first free RAM address 16, got %d", addr)
	}

	addr2 := sut.AddVarEntry("bar")
	if addr2 != 17 {
		t.Errorf("expected next free RAM address 17, got %d", addr2)
	}

	addr3 := sut.AddVarEntry("foo")
	if addr3 != 16 {
		t.Errorf("expected reused address 16 for 'foo', got %d", addr3)
	}
}
