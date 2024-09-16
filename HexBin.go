package mrg

import (
	"strconv"
	"strings"
)

func HexBin(s string) string {
	lines := strings.Split(s, "\n")
	var result string

	for _, line := range lines {
		la := strings.Fields(line)
		n := len(la)

		for i := 0; i < n; i++ {
			if la[i] == "(hex)" || la[i] == "(bin)" {
				if i > 0 {
					var value int64
					var err error

					if la[i] == "(hex)" {
						for a := i - 1; a >= 0; a-- {
							if la[a] != "" {
								value, err = strconv.ParseInt(la[a], 16, 64)
								if err == nil {
									la[a] = strconv.Itoa(int(value))
								}
								break
							}
						}
					} else if la[i] == "(bin)" {
						for v := i - 1; v >= 0; v-- {
							if la[v] != "" {
								value, err = strconv.ParseInt(la[v], 2, 64)
								if err == nil {
									la[v] = strconv.Itoa(int(value))
								}
								break
							}
						}
					}
				}
				la[i] = ""
			}
		}

		lineResult := RemoveEmptyStrings(la)
		result += strings.Join(lineResult, " ") + "\n"
	}

	return result
}
