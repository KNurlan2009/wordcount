package countutil

import "strings"

func WordCounter(in string) int {
	if len(strings.TrimSpace(in)) < 1 {
		return 0
	}
	s := strings.Split(strings.TrimSpace(in), " ")
	return len(s)
}
