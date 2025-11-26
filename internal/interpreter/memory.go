package interpreter

type Memory struct {
	Data []int
}

func NewMemory(size int) *Memory {
	return &Memory{
		Data: make([]int, size),
	}
}

func (m *Memory) Read(addr int) int {
	if addr < 0 || addr >= len(m.Data) {
		return 0
	}
	return m.Data[addr]
}

func (m *Memory) Write(addr int, value int) {
	if addr < 0 || addr >= len(m.Data) {
		return
	}
	m.Data[addr] = value
}
