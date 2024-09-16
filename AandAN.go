package mrg

import (
	"strings"
)

// Aandan processes the input string to modify occurrences of "a" based on the following word while preserving newlines.
func Aandan(s string) string {
	lines := strings.Split(s, "\n")
	var result string

	for _, line := range lines {
		la := strings.Fields(line)
		res := ""

		for i := 0; i < len(la); i++ {
			if strings.ToLower(la[i]) == "a" && i != len(la)-1 {
				if len(la[i+1]) != 0 && (la[i+1][0] == 'A' || la[i+1][0] == 'E' || la[i+1][0] == 'I' || la[i+1][0] == 'O' || la[i+1][0] == 'U' || la[i+1][0] == 'H') {
					res += la[i] + "n"
				} else if len(la[i+1]) != 0 && (la[i+1][0] == 'a' || la[i+1][0] == 'e' || la[i+1][0] == 'i' || la[i+1][0] == 'o' || la[i+1][0] == 'u' || la[i+1][0] == 'h') {
					res += la[i] + "n"
				} else {
					res += la[i]
				}
			} else {
				res += la[i]
			}
			res += " "
		}

		result += strings.TrimSpace(res) + "\n"
	}

	return result
}
