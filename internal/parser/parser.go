package parser

import (
	"errors"
	"strings"
)

type InstructionType int

var (
	ErrUnsupportedInstruction = errors.New("instruction is not supported")
)

const (
	A_INSTRUCTION InstructionType = iota
	C_INSTRUCTION
	L_INSTRUCTION
)

type Parser struct {
	FileName     string
	lines        []string
	currPosition int
	currLine     string
}

func New(fileName string, content []string) *Parser {
	return &Parser{FileName: fileName, lines: content}
}

// ==============================================
// 			Nand2Tetris Parser API:
// ==============================================

func (p *Parser) HasMoreLines() bool {
	return p.currPosition < len(p.lines)
}

func (p *Parser) InstructionType() (InstructionType, error) {
	line := strings.TrimSpace(p.currLine)
	switch {
	case strings.HasPrefix(line, "@"):
		return A_INSTRUCTION, nil
	case strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")"):
		return L_INSTRUCTION, nil
	case strings.ContainsAny(line, ";="):
		return C_INSTRUCTION, nil
	default:
		return 0, ErrUnsupportedInstruction
	}
}

func (p *Parser) Symbol() string {
	line := strings.TrimSpace(p.currLine)
	line = strings.TrimPrefix(line, "@")
	line = strings.TrimPrefix(line, "(")
	return strings.TrimSuffix(line, ")")
}

func (p *Parser) Advance() {
	line := strings.TrimSpace(p.lines[p.currPosition])
	p.currPosition++

	if idx := strings.Index(line, "//"); idx != -1 {
		line = line[:idx]
	}
	p.currLine = strings.TrimSpace(line)
}

func (p *Parser) Dest() string {
	if strings.Contains(p.currLine, "=") {
		parts := strings.SplitN(p.currLine, "=", 2)
		return strings.TrimSpace(parts[0])
	}
	return ""
}

func (p *Parser) Comp() string {
	line := p.currLine
	if strings.Contains(line, "=") {
		parts := strings.SplitN(line, "=", 2)
		line = parts[1]
	}
	if strings.Contains(line, ";") {
		parts := strings.SplitN(line, ";", 2)
		return strings.TrimSpace(parts[0])
	}
	return strings.TrimSpace(line)
}

func (p *Parser) Jump() string {
	if strings.Contains(p.currLine, ";") {
		parts := strings.SplitN(p.currLine, ";", 2)
		return strings.TrimSpace(parts[1])
	}
	return ""
}

func (p *Parser) Reset() {
	p.currLine = ""
	p.currPosition = 0
}
