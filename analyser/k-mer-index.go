package analyser

type IndexArt struct {
	K int
	I []Index
}

type Index struct {
	Key string
	Val int
}

func ConstructIA(t string, k int) IndexArt {
	IA := IndexArt{}

	IA.Init(t, k)
	indexSort(IA.I, 1)

	return IA
}

func (IA *IndexArt) Init(t string, k int) {
	IA.K = k

	for i := 0; i <= len(t)-k; i++ {
		idx := Index{
			Key: t[i : i+k],
			Val: i,
		}
		IA.I = append(IA.I, idx)
	}

}

func indexSort(idx []Index, depth int) {
	for x := range idx {
		y := x + 1
		for y = range idx {
			if idx[x].Key < idx[y].Key {
				idx[x], idx[y] = idx[y], idx[x]
			}
		}
	}
}

func (IA *IndexArt) QueryKMer(p string) []int {
	var hits []int
	kmer := p[:IA.K]

	for i := 0; i < len(IA.I); i++ {
		if kmer != IA.I[i].Key[0:0] {
			break
		}
		if kmer == IA.I[i].Key[0:IA.K] {
			hits = append(hits, IA.I[i].Val)
		}
	}
	return hits
}
