package asmjson

import (
	"encoding/json"
	"fmt"
	"os"

	"practica3/internal/ir"
)

func ParseJSON(path string) (*ir.Program, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s: %w", path, err)
	}

	var jp JSONProgram
	if err := json.Unmarshal(data, &jp); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	prog := ir.NewProgram()

	for _, j := range jp.Instructions {
		raw := ir.RawInstruction{
			Op:   j.Op,
			Args: j.Args,
		}

		instr, err := ir.FromRaw(raw)
		if err != nil {
			return nil, fmt.Errorf("instruction error: %w", err)
		}

		prog.Add(instr)
	}

	return prog, nil
}
