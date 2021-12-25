package util

import "strings"

func JoinString(strs ...string) string {
	var stringBuilder strings.Builder
	for _, str := range strs {
		stringBuilder.WriteString(str)
	}
	return stringBuilder.String()
}
