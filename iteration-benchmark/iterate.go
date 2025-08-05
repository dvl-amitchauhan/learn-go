package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	// Strings in Go are immutable, copying memory - 4 allocs/op
	// var repeated string
	// for i := 0; i < repeatCount; i++ {
	// 	repeated += character
	// }
	// return repeated

	// WriteString minimizes memory copying - 1 allocs/op
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
