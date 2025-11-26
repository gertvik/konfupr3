package interpreter

import (
	"os"

	"practica3/internal/ir"
)

type Interpreter struct {
	Mem      *Memory
	Acc      int // аккумулятор
	PC       int // счётчик команд (в байтах)
	CodeSize int // размер загруженной программы в байтах
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Mem:      NewMemory(65536),
		PC:       0,
		Acc:      0,
		CodeSize: 0,
	}
}

func (it *Interpreter) LoadBinary(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	it.CodeSize = len(data)

	// Загружаем машинный код в память, начиная с адреса 0.
	for i, b := range data {
		it.Mem.Write(i, int(b))
	}

	return nil
}

func (it *Interpreter) Step() bool {
	// Нужны как минимум 2 байта для очередной инструкции
	if it.PC+1 >= it.CodeSize {
		return false
	}

	// Читаем байты инструкции из объединённой памяти
	b1 := byte(it.Mem.Read(it.PC))
	b2 := byte(it.Mem.Read(it.PC + 1))
	it.PC += 2

	instr := DecodeInstruction(b1, b2)

	switch instr.Op {
	case ir.OpLoadConst:
		it.Acc = instr.Args["value"]

	case ir.OpLoadMem:
		addr := it.Acc
		it.Acc = it.Mem.Read(addr)

	case ir.OpStoreMem:
		addr := instr.Args["addr"]
		it.Mem.Write(addr, it.Acc)

	case ir.OpNeg:
		it.Acc = -it.Acc
	}

	return true
}

func (it *Interpreter) Run() {
	for it.Step() {
	}
}
