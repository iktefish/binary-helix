package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/client"
	"github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/server"
	"github.com/iktefish/binary-helix/types"
	"github.com/iktefish/binary-helix/utils"
	"github.com/iktefish/binary-helix/workers"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("Please enter an arg!")
		return
	}

	// if arg[1] == "RegisterNode" {
	// 	client.RegisterNode("172.17.0.3:4042", "binary-helix_c2")
	// }

	if arg[1] == "TotalBasesOfEach" {
		As, Cs, Gs, Ts := analyser.TotalBasesOfEach()

		fmt.Println("As ~~> ", As)
		fmt.Println("Cs ~~> ", Cs)
		fmt.Println("Gs ~~> ", Gs)
		fmt.Println("Ts ~~> ", Ts)
	}

	if arg[1] == "BoyerMoore" {
		// pBM := types.BoyerMoore{}
		t := "GCTACGATCTAGAATCTA"
		p := "TCTA"
		// t := "TCTA"
		// pBM.Init(p)
		pBM := types.ConstructBM(p)

		// p_bm := types.BoyerMoore(p)

		fmt.Println("t[7:11] ~~> ", t[7:11])
		fmt.Println("t[14:18] ~~> ", t[14:18])

		fmt.Println(pBM.Bad_Character_Rule(2, "T"))
		fmt.Println(analyser.BoyerMoore(p, pBM, t))
	}

	if arg[1] == "ExactMatch" {
		t := "GCTACGATCTAGAATCTA"
		p := "TC"
		extMatch := analyser.ExactMatch(p, t)
		fmt.Println(extMatch)
	}

	if arg[1] == "KMer" {
		analyser.ConstructIA("GCTACGATCTAGAATAACTA", 2)
	}

	if arg[1] == "Read" {
		// path := "test/input/phix.fa"
		// path := "./test/input/sra_data.fastq"
		path := "./test/input/small_sra_data.fastq"

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, "quality-scores")
	}

	if arg[1] == "Qual" {
		path := "./test/input/small_sra_data.fastq"

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		analyser.Id_SeqQual(splits)
	}

	if arg[1] == "Server" {
		server.Server()
	}

	if arg[1] == "Client" {
		server.Server()
	}

	if arg[1] == "DB" {
		// utils.Admin_EchoDbContents("nodes_db")
		// utils.Admin_EchoDbContents("slices_db")
		// utils.Admin_EchoDbContents("bench_db")
		// utils.Admin_EchoDbContents("Hello_DB")

		// utils.Admin_DummyInComputeNodes()
		// utils.Admin_DummyInSlices()
		// utils.Admin_DummyInBenchmarks()

		// utils.Admin_EchoDbs()

		// utils.Admin_ClearDbAll("nodes_db")
		// utils.Admin_ClearDbAll("slices_db")
		// utils.Admin_ClearDbAll("bench_db")
		// utils.Admin_ClearDbAll("Hello")
		utils.Admin_ClearDbAll("all")

		schema.Test_TimeToPrim()
	}

	// HERE_IT_STARTS:
	/* Command line args */

	/* Help */
	if strings.ToLower(arg[1]) == "help" {
		fmt.Println("Here's help!!!")
	}

	/* Register Node */
	if strings.ToLower(arg[1]) == "register-node" {
		ip_port := arg[2]
		node_name := arg[3]

		out := false
		if node_name != "" {
			out = client.RegisterNode(ip_port, node_name)
		}

		if out != true {
			fmt.Println("FAIL: Registration failed!")
		}
	}

	/* Check Nodes */
	if strings.ToLower(arg[1]) == "check-nodes" {
		client.CheckServers()
	}

	/* Check Nodes */
	if strings.ToLower(arg[1]) == "admin_clear-db" {
		dbs := [4]string{
			"nodes_db",
			"slices_db",
			"bench_db",
			"all",
		}

		if len(arg) < 3 {
			fmt.Println("FAIL: Please provide which database you want purged!")
			fmt.Println("Possible options are:")
			for i := range dbs {
				fmt.Println(i+1, dbs[i])
			}

            return
		}

		db := arg[2]
		if db == dbs[0] {
			utils.Admin_ClearDbAll("nodes_db")
		}
		if db == dbs[1] {
			utils.Admin_ClearDbAll("slices_db")
		}
		if db == dbs[2] {
			utils.Admin_ClearDbAll("bench_db")
		}
		if db == dbs[3] {
			utils.Admin_ClearDbAll("all")
		}
	}

	/* Boyer-Moore */
	if strings.ToLower(arg[1]) == utils.AnalyserList[0] {
        if len(arg) != 4 {
			fmt.Println("FAIL: Please provide proper list of arguments!")
			fmt.Println("Type `binary-helix help` for more information.")

            return
        }

		path := arg[2]
		p := arg[3]

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[0], p)

		// fmt.Println(lineCount)
		// fmt.Println(path)
		// fmt.Println(p)

		// pBM := types.ConstructBM(p)
		// fmt.Println(pBM.Bad_Character_Rule(2, "T"))
		// fmt.Println(analyser.BoyerMoore(p, pBM, t))
	}
}
