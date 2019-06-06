package string_unpacker

import (
	"testing"
)

var testCases = map[string]string{
	"a4bc2d5e": "aaaabccddddde",
	"abcd":     "abcd",
	"°5€1":     "°°°°°€",
	"45":       "",
	"5oijsd":   "",
}

func TestUnpack(t *testing.T) {
	for inputVal, expected := range testCases {
		if resultVal := Unpack(inputVal); resultVal != expected {
			t.Errorf("Strings must by match: %s - %s", expected, resultVal)
		}
	}
}
