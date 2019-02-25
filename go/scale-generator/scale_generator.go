package scale

import "strings"

var stepSizes = map[byte]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

// Scale generates a diatonic scale given a tonic and list of intervals
func Scale(tonic, interval string) []string {
	var chromatic []string
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		chromatic = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		chromatic = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}

	var tonicIndex int
	for i, n := range chromatic {
		if strings.ToLower(n) == strings.ToLower(tonic) {
			tonicIndex = i
			break
		}
	}

	if interval == "" {
		return append(chromatic[tonicIndex:], chromatic[:tonicIndex]...)
	}

	scale := make([]string, len(interval))
	scale[0] = chromatic[tonicIndex]
	previousIndex := tonicIndex
	for i := 1; i < len(interval); i++ {
		step := stepSizes[interval[i-1]]
		index := (previousIndex + step) % len(chromatic)
		scale[i] = chromatic[index]
		previousIndex = index
	}
	return scale
}
