package config

import "strings"

func GetExcludePaths(strInput string) [] string {
	splits := strings.Split(strInput, ":")
	results := []string{}

	for _, split := range splits {
		if split != "" {
			results = append(results, "/"+split +"/")
		}
	}

	return results
}

