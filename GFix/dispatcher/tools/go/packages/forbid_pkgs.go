package packages

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func is_path_forbind(path string) bool {
	return stringInSlice(path,forbidden_list)
}

var forbidden_list []string