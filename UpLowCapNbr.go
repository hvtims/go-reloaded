package mrg

import (
	"fmt"
	"strconv"
	"strings"
)

func UpLowCapNbr(s string) string {
	lines := strings.Split(s, "\n")
	var result string

	Flags := []string{"(cap,", "(low,", "(up,"}

	for _, line := range lines {
		la := strings.Fields(line)

		for i := 0; i < len(la); i++ {
			for _, Flag := range Flags {
				if strings.HasPrefix(la[i], Flag) && i+1 < len(la) {
					numberStr := strings.Trim(la[i+1], " )")
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						fmt.Println("error handling the number:", err)
						continue
					}

					start := i - number
					if start < 0 {
						start = 0
					}
					if number <= 0 {
						la[i] = ""
						la[i+1] = ""
						break
					}

					count := 0
					for j := i - 1; j >= start && count < number; j-- {
						if la[j] == "" {
							continue
						}

						switch Flag {
						case "(cap,":
							la[j] = Capitalize(la[j])
						case "(up,":
							la[j] = strings.ToUpper(la[j])
						case "(low,":
							la[j] = strings.ToLower(la[j])
						}

						count++
					}

					la[i] = ""
					la[i+1] = ""
					break
				}
			}
		}

		la = RemoveEmptyStrings(la)
		result += strings.Join(la, " ") + "\n"
	}

	return result
}
