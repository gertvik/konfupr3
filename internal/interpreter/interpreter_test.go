package interpreter

import (
	"testing"

	"practica3/internal/emitter"
	"practica3/internal/ir"
)

func encodeProgram(t *testing.T, prog *ir.Program) []byte {
	t.Helper()
	var machineCode []byte
	for _, instr := range prog.Instructions {
		enc, err := emitter.EncodeInstruction(instr)
		if err != nil {
			t.Fatalf("encode failed: %v", err)
		}
		machineCode = append(machineCode, enc[0], enc[1])
	}
	return machineCode
}

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

	it := NewInterpreter()
	it.LoadBytes(encodeProgram(t, prog))
	it.Run()

	if got := it.Mem.Read(64); got != -123 {
		t.Fatalf("memory[64] = %d, want -123", got)
	}
}

func TestVectorUnaryNeg(t *testing.T) {
	values := []int{5, 13, 0, 255, 1024, 4095}
	srcBase := 100
	dstBase := 200

	prog := ir.NewProgram()

	for i, v := range values {
		prog.Add(&ir.Instruction{
			Op:   ir.OpLoadConst,
			Args: map[string]int{"value": v},
		})
		prog.Add(&ir.Instruction{
			Op:   ir.OpStoreMem,
			Args: map[string]int{"addr": srcBase + i},
		})
	}

	for i := range values {
		prog.Add(&ir.Instruction{
			Op:   ir.OpLoadConst,
			Args: map[string]int{"value": srcBase + i},
		})
		prog.Add(&ir.Instruction{Op: ir.OpLoadMem, Args: map[string]int{}})
		prog.Add(&ir.Instruction{Op: ir.OpNeg, Args: map[string]int{}})
		prog.Add(&ir.Instruction{
			Op:   ir.OpStoreMem,
			Args: map[string]int{"addr": dstBase + i},
		})
	}

	it := NewInterpreter()
	it.LoadBytes(encodeProgram(t, prog))
	it.Run()

	for i, v := range values {
		if got := it.Mem.Read(dstBase + i); got != -v {
			t.Fatalf("memory[%d] = %d, want %d", dstBase+i, got, -v)
		}
	}
}
