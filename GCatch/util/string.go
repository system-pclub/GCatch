package util

import "strings"

func SplitStr2Map(strTarget string, strSep string) map[string]bool {
	splits := strings.Split(strTarget,strSep)
	result := make(map[string]bool)
	for _,split := range splits {
		if split != "" {
			result[split] = true
		}
	}
	return result
}
