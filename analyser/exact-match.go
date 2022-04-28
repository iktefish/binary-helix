package analyser

func ExactMatch(p string, t string) []int {
	var occurrences []int

	for i := 0; i <= len(t)-len(p)+1; i++ {
		match := true
		for j := range p {
			if t[i+j] != p[j] {
				match = false
				break
			}
        if match == true {
            occurrences = append(occurrences, i)
        }
		}
	}

	return occurrences
}
