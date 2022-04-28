package analyser

import (
	"fmt"
	"strings"

	"github.com/iktefish/binary-helix/workers"
)

func Id_SeqQual(s []string) {
	path := "./test/input/small_sra_data.fastq"

	fileExt, processed, lineCount := workers.Reader(path)
	splits := workers.Splitter(fileExt, processed, lineCount)

	split := splits[0]

	fmt.Println(split)

	var quals string
	var seqs string

	nlineCount := 0
	for _, line := range strings.Split(strings.TrimRight(split, "\n"), "\n") {
		nlineCount += 1
		if nlineCount%2 == 0 {
			quals += line
		} else {
			seqs += line
		}
	}
	fmt.Println("Q~>", quals)
	fmt.Println("S~>", seqs)

	// get_Q(quals)
	// get_Phred33(get_Q(quals))
	// fmt.Println(get_Phred33(get_Q(quals)))
	fmt.Println(get_Q(quals))
}

func conv_Phred33ToQ(phredRune rune) int {
	return int(phredRune) - 33
}

func get_Q(phred string) []int {
	var Q []int

	for _, phredRune := range phred {
		conv_Phred33ToQ(phredRune)
		Q = append(Q, int(phredRune))
	}

	return Q
}

func conv_QToPhred33(Q int) string {
	return string(Q + 33)
}

func get_Phred33(Qs []int) string {
	var phred string

	for _, Q := range Qs {
		conv_QToPhred33(Q)
		phred += string(Q)
	}

	return phred
}
