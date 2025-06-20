package controller

import (
	"fmt"

	"github.com/sanmoskalenko/hack-assembler/internal/fileio/loader"
	"github.com/sanmoskalenko/hack-assembler/internal/fsm"
)

type Loader interface {
	LoadASM(dir string) ([]loader.File, error)
}

type Assembler interface {
	Assemble(files []loader.File, outDir string) map[string][]byte
}

type Exporter interface {
	Export(file string, outDir string, data []byte) error
}

type Controller struct {
	FSM       *fsm.FSM
	InputDir  string
	OutputDir string
	Loader    Loader
	Assembler Assembler
	Exporter  Exporter
}

func New(
	fsm *fsm.FSM,
	inputDir string,
	outputDir string,
	loader Loader,
	assembler Assembler,
	exporter Exporter,
) *Controller {
	return &Controller{
		FSM:       fsm,
		InputDir:  inputDir,
		OutputDir: outputDir,
		Loader:    loader,
		Assembler: assembler,
		Exporter:  exporter,
	}
}

func (c *Controller) Run() {
	c.FSM.Send(fsm.EventPayload{Event: fsm.EventStart})

	files := []loader.File{}
	result := map[string][]byte{}

	actions := map[fsm.State]func() fsm.EventPayload{
		fsm.Preparing: func() fsm.EventPayload {
			var err error
			files, err = c.Loader.LoadASM(c.InputDir)
			if err != nil {
				return fsm.EventPayload{
					Event:   fsm.EventFail,
					Message: err.Error(),
				}
			}
			return fsm.EventPayload{
				Event:   fsm.EventSuccess,
				Message: fmt.Sprintf("Found %d .asm files", len(files)),
			}
		},
		fsm.Processing: func() fsm.EventPayload {
			result = c.Assembler.Assemble(files, c.OutputDir)
			return fsm.EventPayload{Event: fsm.EventSuccess}
		},
		fsm.Exporting: func() fsm.EventPayload {
			for name, binary := range result {
				if err := c.Exporter.Export(name, c.OutputDir, binary); err != nil {
					return fsm.EventPayload{
						Event:   fsm.EventFail,
						Message: err.Error(),
					}
				}
			}
			return fsm.EventPayload{
				Event:   fsm.EventSuccess,
				Message: fmt.Sprintf("Output written to: %s", c.OutputDir),
			}
		},
	}

	for !c.FSM.IsTerminal() {
		c.FSM.Dispatch(actions)
	}
	c.FSM.Send(fsm.EventPayload{Event: fsm.EventFinish})
}
