package handlers

import "strings"

func Concat(s ...string) *string {
	var b strings.Builder
	for _, el := range s {
		b.WriteString(el)
	}
	var str string = b.String()
	return &str
}
