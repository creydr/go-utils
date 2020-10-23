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

// UniqueStrings returns a slice with unique strings from the given slice
func UniqueStrings(slice []string) []string {
	keys := make(map[string]bool)
	result := make([]string, 0, len(slice))

	for _, val := range slice {
		if _, found := keys[val]; !found {
			result = append(result, val)
			keys[val] = true
		}
	}

	return result
}

// IsEqualStringSlice returns true if the given two slices are equal
func IsEqualStringSlice(s1 []string, s2 []string) bool {
	if &s1 == &s2 {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	for k := range s1 {
		if s1[k] != s2[k] {
			return false
		}
	}

	return true
}

// IsEqualIntSlice returns true if the given two slices are equal
func IsEqualIntSlice(s1 []int, s2 []int) bool {
	if &s1 == &s2 {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	for k := range s1 {
		if s1[k] != s2[k] {
			return false
		}
	}

	return true
}
