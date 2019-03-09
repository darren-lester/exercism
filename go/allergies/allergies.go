package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies calculates the list of substances an allergy score corresponds to
func Allergies(score uint) []string {
	allergies := make([]string, 0)
	for _, allergen := range allergens {
		if score%2 == 1 {
			allergies = append(allergies, allergen)
		}
		score = score >> 1
	}
	return allergies
}

// AllergicTo returns whether or not a given substance is an allergen for a given score
func AllergicTo(score uint, substance string) bool {
	for _, allergen := range allergens {
		if allergen == substance {
			return score%2 == 1
		}
		score = score >> 1
	}
	return false
}
