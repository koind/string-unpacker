package string_unpacker

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	backslash = `\`
)

type unpacker struct {
	symbols []string
}

func NewUnpacker() *unpacker {
	return &unpacker{}
}

func (u *unpacker) Unpack(text string) string {
	if text == "" || isInt(text) {
		return ""
	}

	if r, _ := utf8.DecodeRuneInString(text); unicode.IsNumber(r) {
		return ""
	}

	u.symbols = make([]string, 0, len(text))

	b := []byte(text)

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)

		if unicode.IsNumber(r) {
			u.propagate(r)
		} else {
			u.append(string(r))
		}

		b = b[size:]
	}

	return u.join()
}

func (u *unpacker) append(str string) {
	u.symbols = append(u.symbols, str)
}

func (u *unpacker) propagate(r rune) {
	previousSymbol := u.symbols[len(u.symbols)-1]
	intVal, _ := strconv.Atoi(string(r))

	str := strings.Repeat(previousSymbol, intVal)
	u.symbols[len(u.symbols)-1] = str
}

func (u *unpacker) join() string {
	return strings.Join(u.symbols, "")
}

func isInt(strVal string) bool {
	if _, err := strconv.Atoi(strVal); err == nil {
		return true
	}

	return false
}

func isBackslash(r rune) bool {
	return string(r) == backslash
}
