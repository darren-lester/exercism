package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func freq(s string, ch chan FreqMap) {
	ch <- Frequency(s)
}

func ConcurrentFrequency(strings []string) FreqMap {
	ch := make(chan FreqMap, len(strings))

	for _, str := range strings {
		go freq(str, ch)
	}

	freqMap := FreqMap{}

	for i := 0; i < len(strings); i++ {
		m := <-ch
		for k, v := range m {
			freqMap[k] += v
		}
	}

	return freqMap
}
