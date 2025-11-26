package ir

import "fmt"

type Opcode int

const (
	OpLoadConst Opcode = iota
	OpLoadMem
	OpStoreMem
	OpNeg
	OpLoad Opcode = iota
	OpStore
	OpAdd
	OpSub
	OpMul
	OpJump
	OpHalt
)

func parseOpcode(op string) (Opcode, error) {
	switch op {
	case "LOAD_CONST":
		return OpLoadConst, nil
	case "LOAD_MEM":
		return OpLoadMem, nil
	case "STORE_MEM":
		return OpStoreMem, nil
	case "NEG":
		return OpNeg, nil
	}
	return 0, fmt.Errorf("unknown opcode: %s", op)
}

type Instruction struct {
	Op   Opcode
	Args map[string]int
}

type RawInstruction struct {
	Op   string
	Args map[string]interface{}
}

func FromRaw(r RawInstruction) (*Instruction, error) {
	opcode, err := parseOpcode(r.Op)
	if err != nil {
		return nil, err
	}

	args := make(map[string]int)
	for key, v := range r.Args {
		f, ok := v.(float64)
		if !ok {
			return nil, fmt.Errorf("argument %s must be a number", key)
		}
		args[key] = int(f)
	}

	return &Instruction{
		Op:   opcode,
		Args: args,
	}, nil
}
