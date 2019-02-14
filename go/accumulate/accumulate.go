package accumulate

// Accumulate applies a converter to a collection of strings
func Accumulate(collection []string, converter func(string) string) []string {
	convertedCollection := make([]string, len(collection))
	for i, s := range collection {
		convertedCollection[i] = converter(s)
	}
	return convertedCollection
}
