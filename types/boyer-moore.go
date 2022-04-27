package types

import (
	"log"
	"math"
	"strings"
)

/*
   NOTE:
   Encapsulates pattern and associated Boyer-Moore preprocessing.
*/
type BoyerMoore struct {
	P             string
	Alphabet      string
	AMap          map[string]int
	Bad_Char      [][]int
	Big_l         []int
	Small_l_Prime []int
}

const alphabet = "ACGT"

/*
   NOTE:
   Constructor for returning a BoyerMoore type.
*/
func Construct(p string) BoyerMoore {
	B := BoyerMoore{}
	B.Init(p)

	return B
}

/*
   NOTE:
   This method must be invoked after structure construction.
*/
func (B *BoyerMoore) Init(p string) {
	B.P = p
	B.Alphabet = alphabet

	/* Create map from alphabet characters to ingeters */
	B.AMap = make(map[string]int)
	for i := range B.Alphabet {
		B.AMap[string(B.Alphabet[i])] = i // NOTE: Compiler cannot use `B.alphabet[i]` as index since its a `byte`
	}

	/* Make bad character rule */
	B.Bad_Char = dense_Bad_Char_Tab(p, B.AMap)

	/* Create good suffix table */
	_, B.Big_l, B.Small_l_Prime = good_Suffix_Table(p)
}

/*
   NOTE:
   Return number of skips given by bad character rule at offset i.
*/
func (B *BoyerMoore) Bad_Character_Rule(i int, c string) int {
	/* TODO: Assert that `c` is in `B.amap` */

	ci := B.AMap[c]

	/* ASSERT */
	// if i >= (B.Bad_Char[i][ci] - 1) {
	// 	log.Fatal("FATAL: Corrupt Input!")
	// }

	return i - (B.Bad_Char[i][ci] - 1)
}

/*
   NOTE:
   Given a mismatch at offset i, return amount to shift as determined
   by (weak) good suffix rule.
*/
func (B *BoyerMoore) Good_Suffix_Rule(i int) int {
	length := len(B.Big_l)

	if i >= length {
		log.Fatal("FATAL: Corrupt Input!")
	}

	if i == length-1 {
		return 0
	}

	i += 1 // `i` points to the leftmost matching position of pattern `P`
	if B.Big_l[i] > 0 {
		return length - B.Big_l[i]
	}
	return length - B.Small_l_Prime[i]
}

/*
   NOTE:
   Return amount to shift in case where P matches T
*/
func (B *BoyerMoore) Match_Skip() int {
	return len(B.Small_l_Prime) - B.Small_l_Prime[1]
}

/* NOTE:
   Use Z algorithm (Gusfield theorem 1.4.1) to preprocess s.
*/
func z_Array(s string) []int {

	/* ASSERT */
	// if len(s) <= 1 {
	// 	log.Fatal("FAIL: Wrong length of input!")
	// }

	var z = []int{len(s)}

	for i := range s {
		if i == len(s)-1 {
			break
		}
		z = append(z, 0)
	}

	/* Initial comparison for s[i:] with prefix */
	for i := range s[1:] {
		I := i + 1 // NOTE: i is sill 0 at 1st iteration, so we have to add 1 to be within bounds

		if s[I] == s[I-1] {
			z[1] += 1
		} else {
			break
		}
	}

	r, l := 0, 0
	if z[l] > 0 {
		r, l = z[1], 1
	}

	for k := range s[2:] {
		K := k + 2 // NOTE: k is sill 0 at 1st iteration, so we have to add 2 to be within bounds

		if z[K] != 0 {
			log.Fatal("FAIL: Wrong length of input!")
		}

		if (K) > r {
			/* Case 1 */
			for i := range s[K:] {
				I := i + K // NOTE: i is sill 0 at 1st iteration, so we have to add K to be within bounds

				if s[I] == s[I-K] {
					z[K] += 1
				} else {
					break
				}
			}
			r, l = K+z[K]-1, K
		} else {
			/* CASE 2a: Calculate the length of beta */
			nbeta := r - K + 1
			zkp := z[K-l]
			if nbeta > zkp {
				/* Case 2a: Zkp wins */
				z[K] = zkp
			} else {
				/* CASE 2b: Compare character just past r */
				nmatch := 0
				for i := range s[r+1:] {
					I := i + r + 1

					if s[I] == s[I-K] {
						nmatch += 1
					} else {
						break
					}
					l, r = K, r+nmatch
					z[K] = r - K + 1
				}
			}
		}
	}

	return z
}

/*
   NOTE:
   Compile the N array (Gusfield theorem 2.2.2) from the z array.
*/
func n_Array(s string) []int {
	var rSlice []string
	for _, r := range s {
		rSlice = append([]string{string(r)}, rSlice...)

		/* For visualisation */
		// fmt.Println(s)
		// fmt.Println(reversedS)
		// fmt.Println(strings.Join(rSlice, ""))
	}

	rS := strings.Join(rSlice, "")

	iS := z_Array(rS)

	var riSlice []int
	for _, r := range iS {
		riSlice = append([]int{r}, riSlice...)
	}

	return riSlice
}

/*
   NOTE:
   Compile the N array (Gusfield theorem 2.2.2) using p and N array.
   L'[i] = largest index j less than n such that N[j] = |P[i:]|
*/
func big_l_Prime_Array(p string, n []int) []int {
	var lp []int
	// lp = append(lp, len(p))
	for i := range p {
		if i+1 <= len(p) {
			lp = append(lp, 0)
		}
	}

	for j := range p {
		i := len(n) - n[j]
		if i < len(p) {
			lp[i] = j + 1
		}
	}
	return lp

}

/*
   NOTE:
   Compile L array (Gusfield theorem 2.2.2) using p and L' array.
   L[i] = largest index j less than n such that N[j] >= |P[i:]|
*/
func big_l_Array(p string, lp []int) []int {
	var l []int
	for i := range p {
		if i+1 <= len(p) {
			l = append(l, 0)
		}
	}

	l[1] = lp[1]
	for i := range p[2:] {
		I := i + 2
		lf, lpf := float64(l[I-1]), float64(lp[I])
		l[I] = int((math.Max(lf, lpf)))
	}

	return l
}

/*
   NOTE:
   Compile lp' array (Gusfield theorem 2.2.4) using N array.
*/
func small_l_Prime_Array(n []int) []int {
	var small_lp []int
	for i := range n {
		if i+1 <= len(n) {
			small_lp = append(small_lp, 0)
		}
	}

	for i := range n {
		if n[i] == i+1 {
			small_lp[len(n)-i-1] = i + 1
		}
	}
	for i := len(n) - 2; i <= len(n); i-- {
		if i < 0 {
			break
		}

		if small_lp[i] == 0 {
			small_lp[i] = small_lp[i+1]
		}
	}

	return small_lp
}

/*
   NOTE:
   Return table needed to apply good suffix rule.
*/
func good_Suffix_Table(p string) ([]int, []int, []int) {
	n := n_Array(p)
	lp := big_l_Prime_Array(p, n)

	return lp, big_l_Array(p, lp), small_l_Prime_Array(n)

}

/*
   NOTE:
   Given pattern string and list with ordered alphabet characters, create
   and return a dense bad character table. Table is indexed by offset then
   by character.
*/
func dense_Bad_Char_Tab(p string, amap map[string]int) [][]int {
	// m := make(map[string]int)

	var tab [][]int

	var nxt []int
	for _, i := range amap { // NOTE: `i` here is of type `int` and `_` would have type `string`
		if i+1 <= len(amap) {
			nxt = append(nxt, 0)
		}
	}

	for i, c := range p[0:] {

		/* TODO: Assert that `c` is present in `amap` */

		conduit := make([]int, len(nxt))
		copy(conduit, nxt)

		tab = append(tab, conduit)
		nxt[amap[string(c)]] = i + 1
	}

	return tab
}
