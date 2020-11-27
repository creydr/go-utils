package slice

// IsEqualStringSlice returns true if the given two slices are equal
func IsEqualStringSlice(s1 []string, s2 []string) bool {
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
