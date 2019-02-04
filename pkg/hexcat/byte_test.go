package hexcat

import (
	"testing"
)

func TestIsSpace(t *testing.T) {
	testCase := []struct {
		name   string
		input  byte
		expect bool
	}{
		{name: "1", input: 8, expect: true},
		{name: "2", input: 13, expect: true},
		{name: "3", input: 14, expect: false},
		{name: "4", input: 32, expect: true},
		{name: "5", input: 120, expect: false},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			hexcatByte := newHexcatByte(test.input)
			output := hexcatByte.isSpace()

			if output != test.expect {
				t.Errorf("output %t want %t", output, test.expect)
			}
		})

	}
}
