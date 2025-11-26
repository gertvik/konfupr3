package main

import (
	"practica3/internal/asmjson"
	"practica3/internal/cli"
	"practica3/internal/emitter"
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
		// Печать IR
		printer.PrintProgram(program)

		// Печать бинарного кода
		emitter.PrintBinary(program)
		return
	}

	// Этап 2 — запись бинарного файла
	size, err := emitter.WriteBinary(program, flags.OutPath)
	if err != nil {
		cli.Die(err)
	}

	println("Written bytes:", size)
}
