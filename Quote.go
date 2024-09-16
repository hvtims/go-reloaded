package mrg

import "regexp"

func Quote(s string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	s = re.ReplaceAllString(s, `'${1}'`)
	return s
}
