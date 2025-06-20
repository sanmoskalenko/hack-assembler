package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sanmoskalenko/hack-assembler/internal/assembler"
	"github.com/sanmoskalenko/hack-assembler/internal/controller"
	"github.com/sanmoskalenko/hack-assembler/internal/fileio/exporter"
	"github.com/sanmoskalenko/hack-assembler/internal/fileio/loader"
	"github.com/sanmoskalenko/hack-assembler/internal/fsm"
)

func main() {
	inDir := flag.String("d", "", "Input directory containing .asm files (required)")
	outDir := flag.String("o", "gen", "Output directory for .hack files")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: hackasm -d <inputDir> [-o <outputDir>]\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *inDir == "" {
		fmt.Fprintln(os.Stderr, "Error: -d <inputDir> is required")
		flag.Usage()
		os.Exit(1)
	}

	controller.New(
		fsm.New(),
		*inDir,
		*outDir,
		&loader.Loader{},
		&assembler.Assembler{},
		&exporter.Exporter{},
	).Run()
}
