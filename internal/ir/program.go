package ir

type Program struct {
	Instructions []*Instruction
}

func NewProgram() *Program {
	return &Program{Instructions: []*Instruction{}}
}

func (p *Program) Add(i *Instruction) {
	p.Instructions = append(p.Instructions, i)
}
