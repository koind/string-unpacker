package string_unpacker

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Unpack(text string) string {
	if text == "" || isInt(text) {
		return ""
	}

	if r, _ := utf8.DecodeRuneInString(text); unicode.IsNumber(r) {
		return ""
	}

	strSlice := make([]string, 0, len(text))

	b := []byte(text)

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)

		if unicode.IsNumber(r) {
			previousSymbol := strSlice[len(strSlice)-1]
			intVal, _ := strconv.Atoi(string(r))

			str := strings.Repeat(previousSymbol, intVal)
			strSlice[len(strSlice)-1] = str
		} else {
			strSlice = append(strSlice, string(r))
		}

		b = b[size:]
	}

	return strings.Join(strSlice, "")
}

func isInt(strVal string) bool {
	if _, err := strconv.Atoi(strVal); err == nil {
		return true
	}

	return false
}
