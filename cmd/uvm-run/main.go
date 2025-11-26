package main

import (
	"fmt"
	"os"
	"strconv"

	"practica3/internal/interpreter"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: uvm-run <binary> <csv_out> <start> <end>")
		return
	}

	binPath := os.Args[1]
	csvPath := os.Args[2]
	start, _ := strconv.Atoi(os.Args[3])
	end, _ := strconv.Atoi(os.Args[4])

	it := interpreter.NewInterpreter()

	if err := it.LoadBinary(binPath); err != nil {
		panic(err)
	}

	it.Run()

	if err := interpreter.DumpCSV(it.Mem, start, end, csvPath); err != nil {
		panic(err)
	}
}
