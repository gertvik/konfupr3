package asmjson

type JSONInstruction struct {
	Op   string                 `json:"op"`
	Args map[string]interface{} `json:"args"`
}

type JSONProgram struct {
	Instructions []JSONInstruction `json:"instructions"`
}
