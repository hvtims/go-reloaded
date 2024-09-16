package mrg

import (
	"strings"
)

func Punc(s string) string {
	ss := ";:,.?!"
	res := ""
	for i := 0; i < len(s); i++ {
		if strings.Contains(ss, string(s[i])) {
			if res != "" && res[len(res)-1] == 32 {
				res = res[:len(res)-1]
			}
			res += string(s[i])
			if i+1 != len(s) && s[i+1] != 32 {
				res += " "
			}

			continue
		}
		res += string(s[i])

	}
	return res
}
