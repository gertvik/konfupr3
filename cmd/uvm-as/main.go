package main

import (
	"practica3/internal/asmjson"
	"practica3/internal/cli"
	"practica3/internal/printer"
)

func main() {
	flags, err := cli.Parse()
	if err != nil {
		cli.Die(err)
	}

	program, err := asmjson.ParseJSON(flags.SrcPath)
	if err != nil {
		cli.Die(err)
	}

	if flags.TestMode {
		printer.PrintProgram(program)
		return
	}

	// Этап 2
	// emitter.WriteBinary(program, flags.OutPath)
}
