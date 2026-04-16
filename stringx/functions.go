package stringx

import (
	"fmt"
	"strings"
)

func Extract(prefix string, text string, suffix string) string {
	size := len(text)

	if size < len(prefix)+len(suffix) {
		return ""
	}

	for i := range size {
		hasPrefix := prefix != "" && strings.HasPrefix(text[i:], prefix)
		if prefix == "" || hasPrefix {
			if hasPrefix {
				i += 1
			}

			for j := size; j >= 0; j-- {
				hasSuffix := suffix != "" && strings.HasSuffix(text[i:j], suffix)
				if suffix == "" || hasSuffix {
					if hasSuffix {
						j -= 1
					}

					return text[i:j]
				}
			}
		}
	}

	return ""
}
func ToScreamingSnakeCase(input string) string {
	words := []rune{}
	for i, r := range input {
		if i > 0 && r == ' ' {
			words = append(words, '_')
		} else {
			words = append(words, r)
		}
	}
	return strings.TrimSpace(strings.ToUpper(string(words)))
}

func AsStrings(slice []any) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = fmt.Sprint(v)
	}
	return result
}
func TrimAll(slice []string, cutset string) []string {
	result := make([]string, len(slice))
	for i, s := range slice {
		result[i] = strings.Trim(s, cutset)
	}
	return result
}
func Join(v []any, sep string) string {
	values := AsStrings(v)
	joinedArgs := strings.Join(values, sep)

	return strings.TrimSpace(joinedArgs)
}
