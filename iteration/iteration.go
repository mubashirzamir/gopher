package iteration

import "strings"

// Repeat prints a string multiple times.
//
// s: string to print
//
// n: number of times to print
func Repeat(s string, n int) string {
	var repeated strings.Builder
	for i := 0; s != "" && i < n; i++ {
		repeated.WriteString(s)
	}

	return repeated.String()
}
