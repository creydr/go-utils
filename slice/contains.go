package slice

// ContainsString returns true, if the given string can be found in the given slice of strings
func ContainsString(slice []string, search string) bool {
	for _, val := range slice {
		if val == search {
			return true
		}
	}

	return false
}

// ContainsInt returns true, if the given integer can be found in the given slice of integers
func ContainsInt(slice []int, search int) bool {
	for _, val := range slice {
		if val == search {
			return true
		}
	}

	return false
}
