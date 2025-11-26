package emitter

import (
	"fmt"
	"practica3/internal/ir"
)

func PrintBinary(p *ir.Program) error {
	for i, instr := range p.Instructions {
		encoded, err := EncodeInstruction(instr)
		if err != nil {
			return err
		}

		fmt.Printf("Instr %d: 0x%02X 0x%02X\n",
			i, encoded[0], encoded[1])
	}

	return nil
}
