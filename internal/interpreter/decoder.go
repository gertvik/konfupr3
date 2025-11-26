package interpreter

import (
	"practica3/internal/ir"
)

func DecodeInstruction(b1, b2 byte) *ir.Instruction {
	word := uint16(b1)<<8 | uint16(b2)

	a := byte(word >> 12)
	b := int(word & 0x0FFF)

	raw := ir.RawInstruction{}

	switch a {
	case 5:
		raw.Op = "LOAD_CONST"
		raw.Args = map[string]interface{}{"value": float64(b)}

	case 15:
		raw.Op = "LOAD_MEM"
		raw.Args = map[string]interface{}{}

	case 14:
		raw.Op = "STORE_MEM"
		raw.Args = map[string]interface{}{"addr": float64(b)}

	case 3:
		raw.Op = "NEG"
		raw.Args = map[string]interface{}{}

	default:
		raw.Op = "UNKNOWN"
		raw.Args = map[string]interface{}{}
	}

	instr, _ := ir.FromRaw(raw)
	return instr
}
