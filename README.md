# Hack Assembler

A Go implementation of the Hack Assembler, based on the book [The Elements of Computing Systems](https://www.nand2tetris.org/).

This assembler translates `.asm` programs written in Hack assembly language into `.hack` machine code files.

## Usage

Build the binary:

```bash
make build
```

Run with input and output directories:

```bash
./hackasm -d testdata -o out
```

Or use the provided `Makefile`:

```bash
make run
```

## CLI Flags

- `-d` (required): Path to a single `.asm` file or a directory containing `.asm` files
- `-o` (optional): Output directory for generated `.hack` files (default: `gen`)

Note: When running via `make run`, the flags use default values unless overridden with `in=...` and `out=...`.

## Testing

Run all tests:

```bash
make test
```

## Example

Input file (`Add.asm`):

```asm
@2
D=A
@3
D=D+A
@0
M=D
```

Output file (`Add.hack`):

```text
0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000
```