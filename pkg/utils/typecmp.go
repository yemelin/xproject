package utils

// IsMapsEqual compares two maps, return true if they are equal, false otherwise
func IsMapsEqual(a, b map[string]([]float64)) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || !IsFloat64SlicesEqual(v, w) {
			return false
		}
	}

	return true
}

// IsStringSlicesEqual compares two slices, return true if they are equal, false otherwise
func IsStringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, item := range a {
		if b[i] != item {
			return false
		}
	}

	return true
}

// IsFloat64SlicesEqual compares two slices, return true if they are equal, false otherwise
func IsFloat64SlicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, item := range a {
		if b[i] != item {
			return false
		}
	}

	return true
}
