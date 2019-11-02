package wordfilter


func containsFilteredWords(s *string) bool {
	filteredWords := []string{
		"fuck",
		"shit",
		"bitch",
	}
	for _, word := range filteredWords {
		if strings.Contains(*s, word) {
			return true
		}
	}
	return false
}