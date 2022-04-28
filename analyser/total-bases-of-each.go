package analyser

import (
	"github.com/iktefish/binary-helix/utils"
	// "fmt"
	"io/ioutil"
)

func TotalBasesOfEach(path string) (int, int, int, int) {
	genome, err := ioutil.ReadFile(path)
	if err != nil {
        utils.HandleError(err)
	}

    // fmt.Println("gnome ~~> ", string(genome))

	counts := make(map[string]int)

	counts["A"] = 0
	counts["C"] = 0
	counts["G"] = 0
	counts["T"] = 0

    for _, base := range genome {
        counts[string(base)] += 1
    }

    As := counts["A"]
    Cs := counts["C"]
    Gs := counts["G"]
    Ts := counts["T"]

    return As, Cs, Gs, Ts
}
