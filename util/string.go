package util

import "strings"

func Divide_str(names_str string) map[string]bool {
	splits := strings.Split(names_str,":")
	result := make(map[string]bool)
	for _,split := range splits {
		if split != "" {
			result[split] = true
		}
	}
	return result
}
