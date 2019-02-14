package hexcat

import (
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	"strconv"
)

type Options struct {
	Colors bool
	Bytes  int
	File   os.File
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run(path string, options *Options) []HexcatLine {
	file, err := os.Open(path)
	check(err)

	stat, err := file.Stat()
	check(err)

	fileSize := stat.Size()
	bytes, err := ioutil.ReadFile(path)
	check(err)

	var lines []HexcatLine
	hexcatLine := HexcatLine{}

	for seekIndex, element := range bytes {
		hexcatByte := newHexcatByte(element)

		// Last byte
		if seekIndex == int(fileSize)-1 {
			hexcatLine.appendByte(hexcatByte)
			lines = append(lines, hexcatLine)
		} else {
			if seekIndex > 0 && seekIndex%options.Bytes == 0 {
				lines = append(lines, hexcatLine)
				hexcatLine = HexcatLine{}
			}
			hexcatLine.appendByte(hexcatByte)
		}
	}

	return lines
}

func Render(lines []HexcatLine, options *Options) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	var line int

	for _, element := range lines {
		table.Append([]string{string(strconv.FormatInt(int64(line), 16)), element.toHexString(), element.toHumanString()})
		line = line + options.Bytes
	}

	table.Render()
}
