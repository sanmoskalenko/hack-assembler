package parser_test

import (
	"testing"

	"github.com/sanmoskalenko/hack-assembler/internal/parser"
)

func TestParserAdvancedAndHasMoreLines(t *testing.T) {
	lines := []string{"@2", "D=A", "(LOOP)", "D;JGT"}
	sut := parser.New("file.asm", lines)

	var count int
	for sut.HasMoreLines() {
		sut.Advance()
		count++
	}
	if count != len(lines) {
		t.Errorf("Expected %d lines, got %d", len(lines), count)
	}
}

func TestInstructionType(t *testing.T) {
	tests := []struct {
		line     string
		expected parser.InstructionType
	}{
		{"@100", parser.A_INSTRUCTION},
		{"(LOOP)", parser.L_INSTRUCTION},
		{"D=A", parser.C_INSTRUCTION},
		{"D;JGT", parser.C_INSTRUCTION},
	}

	for _, tt := range tests {
		sut := parser.New("test.asm", []string{tt.line})
		sut.Advance()
		itype, err := sut.InstructionType()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if itype != tt.expected {
			t.Errorf("For line %q, expected %v, got %v", tt.line, tt.expected, itype)
		}
	}
}

func TestSymbolExtraction(t *testing.T) {
	tests := map[string]string{
		"@10":     "10",
		"@LOOP":   "LOOP",
		"(START)": "START",
	}
	for line, expected := range tests {
		sut := parser.New("test.asm", []string{line})
		sut.Advance()
		sym := sut.Symbol()
		if sym != expected {
			t.Errorf("Expected symbol %q, got %q", expected, sym)
		}
	}
}

func TestDestCompJump(t *testing.T) {
	sut := parser.New("test.asm", []string{"MD=D+1;JGT"})
	sut.Advance()

	if got := sut.Dest(); got != "MD" {
		t.Errorf("Expected Dest to be 'MD', got %q", got)
	}
	if got := sut.Comp(); got != "D+1" {
		t.Errorf("Expected Comp to be 'D+1', got %q", got)
	}
	if got := sut.Jump(); got != "JGT" {
		t.Errorf("Expected Jump to be 'JGT', got %q", got)
	}
}

func TestInvalidInstruction(t *testing.T) {
	sut := parser.New("test.asm", []string{"XYZ"})
	sut.Advance()
	_, err := sut.InstructionType()
	if err == nil {
		t.Error("Expected error for invalid instruction, got nil")
	}
}
