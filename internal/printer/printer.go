package printer

import (
	"fmt"

	"practica3/internal/ir"
)

func PrintProgram(p *ir.Program) {
	for i, instr := range p.Instructions {
		fmt.Printf("Instruction %d:\n", i)
		fmt.Printf("  Opcode: %v\n", instr.Op)
		for k, v := range instr.Args {
			fmt.Printf("  %s: %d\n", k, v)
		}
		fmt.Println()
	}
}
