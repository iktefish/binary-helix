package main

import (
	"fmt"
	"os"

	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/nodes"
	"github.com/iktefish/binary-helix/workers"
)

func main() {
	arg := os.Args
	if arg[1] == "TotalBasesOfEach" {
		As, Cs, Gs, Ts := analyser.TotalBasesOfEach()

		fmt.Println("As ~~> ", As)
		fmt.Println("Cs ~~> ", Cs)
		fmt.Println("Gs ~~> ", Gs)
		fmt.Println("Ts ~~> ", Ts)
	}

	if arg[1] == "Read" {
		workers.Reader("test/input/phix.fa")
	}

    if arg[1] == "Server" {
        nodes.Server()
    }

    // if arg[1] == "Client" {
    //     nodes.Client()
    // }
}
