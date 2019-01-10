package twofer

func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
}
