package interpreter

import (
	"testing"

	"practica3/internal/emitter"
	"practica3/internal/ir"
)

func TestUnaryNegStoresResultInMemory(t *testing.T) {
	prog := ir.NewProgram()
	prog.Add(&ir.Instruction{
		Op:   ir.OpLoadConst,
		Args: map[string]int{"value": 123},
	})
	prog.Add(&ir.Instruction{Op: ir.OpNeg, Args: map[string]int{}})
	prog.Add(&ir.Instruction{
		Op:   ir.OpStoreMem,
		Args: map[string]int{"addr": 64},
	})

	var machineCode []byte
	for _, instr := range prog.Instructions {
		enc, err := emitter.EncodeInstruction(instr)
		if err != nil {
			t.Fatalf("encode failed: %v", err)
		}
		machineCode = append(machineCode, enc[0], enc[1])
	}

	it := NewInterpreter()
	it.LoadBytes(machineCode)
	it.Run()

	if got := it.Mem.Read(64); got != -123 {
		t.Fatalf("memory[64] = %d, want -123", got)
	}
}
