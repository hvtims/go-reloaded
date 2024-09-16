package mrg

import (
	"strings"
	"unicode"
)

func UpLowCap(s string) string {
	lines := strings.Split(s, "\n")
	var result string

	for _, line := range lines {
		la := strings.Fields(line)

		for i, r := range la {
			if r == "(up)" {
				if i != 0 {
					for a := i - 1; a >= 0; a-- {
						if la[a] != "" {
							la[a] = strings.ToUpper(la[a])
							break
						}
					}
				}
				la[i] = ""
			} else if r == "(low)" {
				if i != 0 {
					for a := i - 1; a >= 0; a-- {
						if la[a] != "" {
							la[a] = strings.ToLower(la[a])
							break
						}
					}
				}
				la[i] = ""
			} else if r == "(cap)" {
				if i != 0 {
					for a := i - 1; a >= 0; a-- {
						if la[a] != "" {
							la[a] = Capitalize(la[a])
							break
						}
					}
				}
				la[i] = ""
			}
		}

		la = RemoveEmptyStrings(la)
		result += strings.Join(la, " ") + "\n"
	}

	return result
}

func Capitalize(la string) string {
	s := strings.ToLower(la)
	r := []rune(s)
	for i, v := range r {
		if !unicode.IsLetter(v) {
			continue
		}
		r[i] = unicode.ToUpper(v)
		break
	}
	return string(r)
}

func RemoveEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
