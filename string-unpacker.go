package string_unpacker

import (
	"strconv"
	"strings"
)

func Unpack(text string) string {
	if text == "" || isInt(text) {
		return ""
	}

	strSlice := make([]string, 0, len(text))

	for i := 0; i < len(text); i++ {
		b := text[i]
		symbol := string(b)

		if intVal, err := strconv.Atoi(symbol); err == nil && i != 0 {
			previousSymbol := string(text[i-1])

			str := strings.Repeat(previousSymbol, intVal)
			strSlice[len(strSlice)-1] = str

			continue
		} else {
			strSlice = append(strSlice, symbol)
		}
	}

	return strings.Join(strSlice, "")
}

func isInt(strVal string) bool {
	if _, err := strconv.Atoi(strVal); err == nil {
		return true
	}

	return false
}
