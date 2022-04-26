package utils

func Verify_AnalysisName(s string) bool {
	var analysisPresent = [5]string{
		"longest-common-prefix",
		"total-bases-of-each",
		"complement",
		"reverse-complement",
		"quality-scores",
	}

	for _, a := range analysisPresent {
		if s == a {
			return true
		}
	}

	return false
}
