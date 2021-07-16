package iteration

import "strings"

func Repeat(char byte, times int) string {
	var sb strings.Builder
	for i := 0; i < times; i++ {
		sb.WriteByte(char)
	}
	return sb.String()
}
