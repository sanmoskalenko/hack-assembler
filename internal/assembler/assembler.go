package assembler

import (
	"fmt"

	"github.com/sanmoskalenko/hack-assembler/internal/code"
	"github.com/sanmoskalenko/hack-assembler/internal/fileio/loader"
	"github.com/sanmoskalenko/hack-assembler/internal/parser"
	"github.com/sanmoskalenko/hack-assembler/internal/symboltable"
)

type Assembler struct{}

func (r *Assembler) Assemble(files []loader.File, outDir string) map[string][]byte {
	return Assemble(files, outDir)
}

// Assemble processing pipeline
func Assemble(files []loader.File, outDir string) map[string][]byte {
	result := make(map[string][]byte)
	for _, file := range files {
		p := parser.New(file.Name, file.Data)
		st := symboltable.New()
		c := code.NewCode()
		buildSymbolsTable(p, st)
		p.Reset()

		result[p.FileName] = translateInstructions(p, st, c)
	}
	return result
}

func buildSymbolsTable(p *parser.Parser, st *symboltable.SymbolTable) {
	lineAddress := 0
	for p.HasMoreLines() {
		p.Advance()
		itype, _ := p.InstructionType()
		if itype == parser.L_INSTRUCTION {
			st.AddEntry(p.Symbol(), lineAddress)
		} else {
			lineAddress++
		}
	}
}

func translateInstructions(
	p *parser.Parser,
	st *symboltable.SymbolTable,
	c *code.Code,
) []byte {
	var result []byte
	for p.HasMoreLines() {
		p.Advance()
		switch itype, _ := p.InstructionType(); itype {
		case parser.C_INSTRUCTION:
			{
				comp := c.CompByte(p.Comp())
				dest := c.DestByte(p.Dest())
				jump := c.JumpByte(p.Jump())
				binary := c.ComputeCInstruction(dest, comp, jump)
				result = append(result, fmt.Sprintf("%016b\n", binary)...)
			}
		case parser.A_INSTRUCTION:
			{
				symbol := p.Symbol()
				address := st.GetAddress(symbol)
				binary := fmt.Sprintf("%016b\n", address)
				result = append(result, []byte(binary)...)
			}
		default:
			continue
		}
	}
	return result
}
