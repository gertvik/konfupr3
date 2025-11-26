package interpreter

import (
	"encoding/csv"
	"os"
	"strconv"
)

func DumpCSV(mem *Memory, start, end int, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for addr := start; addr <= end; addr++ {
		w.Write([]string{
			strconv.Itoa(addr),
			strconv.Itoa(mem.Read(addr)),
		})
	}

	return nil
}
