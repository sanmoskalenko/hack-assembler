package code_test

import (
	"testing"

	"github.com/sanmoskalenko/hack-assembler/internal/code"
)

func TestPredefinedInstruction(t *testing.T) {
	sut := code.NewCode()
	jumpTest := map[string]byte{
		"JGT": 0b000,
		"JEQ": 0b010,
		"JGE": 0b011,
	}
	destTest := map[string]byte{
		"M":  0b001,
		"D":  0b010,
		"MD": 0b011,
		"A":  0b100,
		"AM": 0b101,
	}
	compTest := map[string]byte{
		"0":  0b0101010,
		"1":  0b0111111,
		"-1": 0b0111010,
		"D":  0b0001100,
		"A":  0b0110000,
	}

	if got := len(sut.Comp); got != 28 {
		t.Errorf("expected 28 comp symbols, got: %d", got)
	}
	if got := len(sut.Dest); got != 7 {
		t.Errorf("expected 7 dest symbols, got: %d", got)
	}
	if got := len(sut.Jump); got != 7 {
		t.Errorf("expected 7 jump symbols, got: %d", got)
	}

	for c, want := range jumpTest {
		if got := sut.JumpByte(c); got != want {
			t.Errorf("expected %v bytes for command: %s,  got %v", want, c, got)
		}
	}
	for c, want := range compTest {
		if got := sut.CompByte(c); got != want {
			t.Errorf("expected %v bytes for command: %s,  got %v", want, c, got)
		}
	}
	for c, want := range destTest {
		if got := sut.DestByte(c); got != want {
			t.Errorf("expected %v bytes for command: %s,  got %v", want, c, got)
		}
	}
}

func TestComputeInstructions(t *testing.T) {
	sut := code.NewCode()

	got := sut.ComputeCInstruction(
		sut.DestByte("M"),
		sut.CompByte("D+1"),
		sut.JumpByte("JMP"))
	want := 0b1110011111001111
	if got != uint16(want) {
		t.Errorf("expected %b bytes for command: M=D+1;JMP, got %b", want, got)
	}

	got = sut.ComputeCInstruction(
		sut.DestByte("MD"),
		sut.CompByte("0"),
		sut.JumpByte(""))
	want = 0b1110101010011000
	if got != uint16(want) {
		t.Errorf("expected %b bytes for command: MD=0, got %b", want, got)
	}

	got = sut.ComputeCInstruction(
		sut.DestByte("A"),
		sut.CompByte("-1"),
		sut.JumpByte(""))
	want = 0b1110111010100000
	if got != uint16(want) {
		t.Errorf("expected %b bytes for command: A=-1, got %b", want, got)
	}

	got = sut.ComputeCInstruction(
		sut.DestByte(""),
		sut.CompByte("!M"),
		sut.JumpByte("JEQ"))
	want = 0b1111110001000010
	if got != uint16(want) {
		t.Errorf("expected %b bytes for command: !M;JEQ, got %b", want, got)
	}
}
