package analyser

import (
	"math"

	"github.com/iktefish/binary-helix/types"
)

func BoyerMoore(p string, pBM types.BoyerMoore, t string) []int {
	i := 0
	var occurences []int

	for i < len(t)-len(p)+1 {
		shift := 1
		mismatched := false
		for j := len(p) - 1; j <= len(p); j-- {
			if j < 0 {
				break
			}

			if p[j] != t[i+j] {
				skip_bc := pBM.Bad_Character_Rule(j, string(t[i+j]))
				skip_gs := pBM.Good_Suffix_Rule(j)
				shift = int(math.Max(float64(shift), (math.Max(float64(skip_bc), float64(skip_gs)))))
				mismatched = true

				break
			}
		}
		if mismatched != true {

			occurences = append(occurences, i)

			skip_gs := pBM.Match_Skip()
			shift = int(math.Max(float64(shift), float64(skip_gs)))
		}
		i += shift
	}

	return occurences
}
