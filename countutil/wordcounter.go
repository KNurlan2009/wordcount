package countutil

import "strings"

func WordCounter(in string) int {
	s := strings.Split(strings.TrimSpace(in), " ")
	return len(s)
}
