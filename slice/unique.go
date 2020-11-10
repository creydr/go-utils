package slice

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
