package ir

import "fmt"

type Opcode int

const (
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
	case "LOAD":
		return OpLoad, nil
	case "STORE":
		return OpStore, nil
	case "ADD":
		return OpAdd, nil
	case "SUB":
		return OpSub, nil
	case "MUL":
		return OpMul, nil
	case "JUMP":
		return OpJump, nil
	case "HALT":
		return OpHalt, nil
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
