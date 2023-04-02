package stringutil

import "unicode"

func ToUpper(input string) string {

	r := []rune(input)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r)
}
