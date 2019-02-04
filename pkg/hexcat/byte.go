package hexcat

import (
	"encoding/hex"
	"strings"
)

type HexcatByte struct {
	value    int    // Decimal representation
	hexValue string // Hexadecimal representation as a string
}

func newHexcatByte(seekedByte byte) *HexcatByte {
	hexcatByte := HexcatByte{}
	hexcatByte.value = int(seekedByte)
	hexcatByte.hexValue = hex.EncodeToString([]byte{seekedByte})

	return &hexcatByte
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func (b *HexcatByte) isSpace() bool {
	var SPACE = []int{8, 9, 10, 11, 12, 13, 32}

	return contains(SPACE, b.value)
}

func (b *HexcatByte) isPrintable() bool {
	var PRINTABLE_CHAR = []int{33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126}

	return contains(PRINTABLE_CHAR, b.value)
}

func (b *HexcatByte) isControlCode() bool {
	var CONTROL_CODE = []int{0, 1, 2, 3, 4, 5, 6, 7, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

	return contains(CONTROL_CODE, b.value)
}

func (b *HexcatByte) colorize() string {
	if b.isControlCode() {
		return "\033[1;30m"
	} else if b.isSpace() {
		return "\033[0;31m"
	} else if b.isPrintable() {
		return "\033[0;32m"
	} else {
		return "\033[0;33m"
	}

	return ""
}

func (b *HexcatByte) print() string {
	str := string(b.value)

	if b.isControlCode() {
		str = "•"
	} else if b.isSpace() || b.isPrintable() {
		m := make(map[string]string)
		m["\b"] = "_"
		m["\t"] = "_"
		m["\n"] = "_"
		m["\v"] = "_"
		m["\f"] = "_"
		m["\r"] = "_"

		for k, v := range m {
			str = strings.Replace(str, k, v, -1)
		}

	} else {
		str = "×"
	}

	return str
}
