package symboltable

import (
	"strconv"
)

type SymbolTable struct {
	symbols     map[string]int
	nextRamAddr int
}

func New() *SymbolTable {
	table := SymbolTable{
		nextRamAddr: 16,
		symbols: map[string]int{
			"SP":     0,
			"LCL":    1,
			"ARG":    2,
			"THIS":   3,
			"THAT":   4,
			"SCREEN": 16384,
			"KBD":    24576,
		},
	}
	for i := range 16 {
		table.AddEntry("R"+strconv.Itoa(i), i)
	}
	return &table
}

func (t *SymbolTable) Contains(s string) bool {
	_, ok := t.symbols[s]
	return ok
}

func (t *SymbolTable) GetAddress(s string) int {
	return t.symbols[s]
}

func (t *SymbolTable) AddVarEntry(s string) int {
	if !t.Contains(s) {
		t.AddEntry(s, t.nextRamAddr)
		t.nextRamAddr++
	}
	return t.GetAddress(s)
}

func (t *SymbolTable) AddEntry(s string, address int) {
	t.symbols[s] = address
}
