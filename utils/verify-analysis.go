package utils

var AnalyserList = [8]string{
	"boyer-moore",
	"complement",
	"reverse-complement",
	"exact-match",
	"k-mer-index",
	"longest-common-prefix",
	"quality-scores",
	"total-bases-of-each",
}

func Check_AnalyzerList(an string) bool {
	for _, a := range AnalyserList {
		if an == a {
			return true
		}
	}

	return false
}
