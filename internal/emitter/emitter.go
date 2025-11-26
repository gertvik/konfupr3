package emitter

import (
	"fmt"
	"os"

	"practica3/internal/ir"
)

// EncodeInstruction — преобразует IR-инструкцию в два байта машинного кода
func EncodeInstruction(instr *ir.Instruction) ([2]byte, error) {
	var a byte
	var b uint16 = 0

	switch instr.Op {
	case ir.OpLoadConst:
		a = 5
		b = uint16(instr.Args["value"])

	case ir.OpLoadMem:
		a = 15

	case ir.OpStoreMem:
		a = 14
		b = uint16(instr.Args["addr"])

	case ir.OpNeg:
		a = 3

	default:
		return [2]byte{}, fmt.Errorf("unknown opcode: %v", instr.Op)
	}

	word := (uint16(a) << 12) | (b & 0x0FFF)

	return [2]byte{
		byte(word >> 8),
		byte(word & 0xFF),
	}, nil
}

// WriteBinary — записывает машинный код в файл
func WriteBinary(p *ir.Program, path string) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	count := 0

	for _, instr := range p.Instructions {
		encoded, err := EncodeInstruction(instr)
		if err != nil {
			return 0, err
		}

		n, err := f.Write(encoded[:])
		if err != nil {
			return 0, err
		}
		count += n
	}

	return count, nil
}
