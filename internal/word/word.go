package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string  {
	return strings.ToLower(s)
}

func UnderscoreToUpperCame(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	return strings.ReplaceAll(s, " ", "")
}

func UnderscoreToLowerCame(s string) string {
	s = UnderscoreToUpperCame(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CameToUnderscore(s string) string {
	var res []rune
	for i, c := range s{
		if i == 0 {
			res = append(res, unicode.ToLower(c))
			continue
		}
		if unicode.IsUpper(c) {
			res = append(res, '_')
		}
		res = append(res, unicode.ToLower(c))
	}
	return string(res)
}