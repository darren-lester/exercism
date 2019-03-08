package etl

import "strings"

// Transform converts a legacy score system to a new score system
func Transform(inputScoreSystem map[int][]string) map[string]int {
  outputScoreSystem := make(map[string]int)

  for score, letters := range inputScoreSystem {
	for _, letter := range letters {
		normalisedLetter := strings.ToLower(letter)
		outputScoreSystem[normalisedLetter] = score
	}
  }
  
  return outputScoreSystem
}