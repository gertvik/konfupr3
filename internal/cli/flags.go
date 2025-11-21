package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	SrcPath  string
	OutPath  string
	TestMode bool
}

func Parse() (*Flags, error) {
	src := flag.String("src", "", "Path to JSON source file")
	out := flag.String("out", "", "Path to output binary file")
	test := flag.Bool("test", false, "Test mode: print IR")

	flag.Parse()

	f := &Flags{
		SrcPath:  *src,
		OutPath:  *out,
		TestMode: *test,
	}

	if f.SrcPath == "" {
		return nil, errors.New("missing required argument: --src file.json")
	}
	if !f.TestMode && f.OutPath == "" {
		return nil, errors.New("missing required argument: --out file.bin")
	}

	return f, nil
}

func Die(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}
