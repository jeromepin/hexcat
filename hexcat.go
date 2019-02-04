package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/jeromepin/hexcat/pkg/hexcat"
	"os"
)

func main() {
	parser := argparse.NewParser("Hexcat", "Make a hexdump with syntaxic coloration")

	colors := parser.Flag("c", "colors", &argparse.Options{Required: false, Help: "Whether to print with colors", Default: true})
	bytesPerLine := parser.Int("b", "bytes", &argparse.Options{Required: false, Help: "Bytes per lines", Default: 16})
	file := parser.File("f", "file", os.O_RDONLY, 0600, &argparse.Options{Required: false, Help: "File"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	opts := &hexcat.Options{
		Colors: *colors,
		Bytes:  *bytesPerLine,
		File:   *file,
	}

	lines := hexcat.Run(file.Name(), opts)
	hexcat.Render(lines, opts)
}
