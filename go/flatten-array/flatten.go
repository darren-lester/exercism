package flatten

// Flatten flattens
func Flatten(element interface{}) []interface{} {
	result := []interface{}{}

	if element == nil {
		return result
	}

	switch value := element.(type) {
	case []interface{}:
		for _, e := range value {
			result = append(result, Flatten(e)...)
		}
	default:
		result = append(result, value)
	}
	return result
}
