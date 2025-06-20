package code

const nulls = 0b000

type Instruction map[string]byte

type Code struct {
	Dest Instruction
	Comp Instruction
	Jump Instruction
}

func NewCode() *Code {
	return &Code{
		Dest: dest,
		Comp: comp,
		Jump: jump,
	}
}

func (c *Code) ComputeCInstruction(dest, comp, jump byte) uint16 {
	var instr uint16 = 0b111 << 13

	instr |= uint16(comp) << 6
	instr |= uint16(dest) << 3
	instr |= uint16(jump)

	return instr
}

func (c *Code) JumpByte(s string) byte {
	if v, ok := c.Jump[s]; ok {
		return v
	}
	return nulls
}

func (c *Code) DestByte(s string) byte {
	if v, ok := c.Dest[s]; ok {
		return v
	}
	return nulls
}

func (c *Code) CompByte(s string) byte {
	if v, ok := c.Comp[s]; ok {
		return v
	}
	return 0
}

var jump = Instruction{
	"JGT": 0b000,
	"JEQ": 0b010,
	"JGE": 0b011,
	"JLT": 0b100,
	"JNE": 0b101,
	"JLE": 0b110,
	"JMP": 0b111,
}

var dest = Instruction{
	"M":   0b001,
	"D":   0b010,
	"MD":  0b011,
	"A":   0b100,
	"AM":  0b101,
	"AD":  0b110,
	"AMD": 0b111,
}

var comp = Instruction{
	"0":   0b0101010,
	"1":   0b0111111,
	"-1":  0b0111010,
	"D":   0b0001100,
	"A":   0b0110000,
	"!D":  0b0001101,
	"!A":  0b0110001,
	"-D":  0b0001111,
	"-A":  0b0110011,
	"D+1": 0b0011111,
	"A+1": 0b0110111,
	"D-1": 0b0001110,
	"A-1": 0b0110010,
	"D+A": 0b0000010,
	"D-A": 0b0010011,
	"A-D": 0b0000111,
	"D&A": 0b0000000,
	"D|A": 0b0010101,
	"M":   0b1110000,
	"!M":  0b1110001,
	"-M":  0b1110011,
	"M+1": 0b1110111,
	"M-1": 0b1110010,
	"D+M": 0b1000010,
	"D-M": 0b1010011,
	"M-D": 0b1000111,
	"D&M": 0b1000000,
	"D|M": 0b1010101,
}
