package hexcat

import (
	"strings"
)

type HexcatLine struct {
	bytes []*HexcatByte
}

func (line *HexcatLine) appendByte(hexcatByte *HexcatByte) {
	line.bytes = append(line.bytes, hexcatByte)
}

func (line *HexcatLine) toHexString() string {
	var bytes []string
	for index := 0; index < len(line.bytes); index++ {
		bytes = append(bytes, line.bytes[index].colorize()+line.bytes[index].hexValue)
	}

	bytes = append(bytes, "\033[0m")

	return strings.Join(bytes, " ")
}

func (line *HexcatLine) toHumanString() string {
	var chars []string

	for index := 0; index < len(line.bytes); index++ {
		b := line.bytes[index]
		chars = append(chars, b.colorize()+b.print())
	}

	chars = append(chars, "\033[0m")

	return strings.Join(chars, "")
}
