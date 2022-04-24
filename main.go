package main

import (
	"fmt"
	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/server"
	"github.com/iktefish/binary-helix/utils"
	"github.com/iktefish/binary-helix/workers"
	"os"
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
		server.Server()
	}

	if arg[1] == "Client" {
		server.Server()
	}

	if arg[1] == "DB" {
		utils.Admin_EchoDbContents("nodes_db")
		utils.Admin_EchoDbContents("slices_db")
		utils.Admin_EchoDbContents("bench_db")
		utils.Admin_EchoDbContents("Hello_DB")

		utils.Admin_DummyInComputeNodes()
		utils.Admin_DummyInSlices()
		utils.Admin_DummyInBenchmarks()

		utils.Admin_EchoDbs()

		if utils.CheckNodeDup() == true {
			fmt.Println("FAIL: An individual host cannot register as more then 1 node!")
		} else {
			fmt.Println("SUCCESS: No duplicate found!")
		}

		schema.Test_TimeToPrim()
	}

}
