package string_unpacker

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type unpacker struct {
	slice []string
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

	u.slice = make([]string, 0, len(text))

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
	u.slice = append(u.slice, str)
}

func (u *unpacker) propagate(r rune) {
	previousSymbol := u.slice[len(u.slice)-1]
	intVal, _ := strconv.Atoi(string(r))

	str := strings.Repeat(previousSymbol, intVal)
	u.slice[len(u.slice)-1] = str
}

func (u *unpacker) join() string {
	return strings.Join(u.slice, "")
}

func isInt(strVal string) bool {
	if _, err := strconv.Atoi(strVal); err == nil {
		return true
	}

	return false
}

func isReverseSolidus(r rune) bool {
	return `\` == string(r)
}
