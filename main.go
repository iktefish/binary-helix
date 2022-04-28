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

		// path := "./test/input/small_sra_data.fastq"

		// fileExt, processed, lineCount := workers.Reader(path)
		// splits := workers.Splitter(fileExt, processed, lineCount)
		// workers.Carrier(splits, "quality-scores")
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
        fmt.Println("")
        fmt.Println("INTRODUCTION:")
        fmt.Println("")
        fmt.Println("\t Welcome to Binary Helix, a distributed Genome analysis system powered by SliceCoin.")
        fmt.Println("")
        fmt.Println("\t Using this system, you can allow scientists, researchers and engineers from all over")
        fmt.Println("\t the world to use a tiny fraction of you smartphone/desktop/leptops' computational power")
        fmt.Println("\t for the purpose of analysing DNA sequences. Doing so you can be a part of may glorious")
        fmt.Println("\t individuals who contribute to seeking, and helping others seek, a greater understanding of")
        fmt.Println("\t things such as Cancer, Down's Syndrome, Aging, Genetic Psychiatric Conditions, Evolution,")
        fmt.Println("\t Language, etc.")
        fmt.Println("")
        fmt.Println("\t You will of-course be paid for donating your computation to the service of science.")
        fmt.Println("")
        fmt.Println("COMMANDS:")
        fmt.Println("")
        fmt.Println("\t binary-helix help\t\t\t\t\t--> Outputs information on how to use this CLI")
        fmt.Println("\t binary-helix register-node IP:PORT\t\t\t--> Registers your device as a server-node to donate computation.")
        fmt.Println("\t binary-helix check-nodes\t\t\t\t--> Outputs the Complement of the DNA of an input .fa file.")
        fmt.Println("\t binary-helix complement FILE\t\t\t\t--> Outputs the Complement of the DNA of an input .fa file.")
        fmt.Println("\t binary-helix reverse-complement FILE\t\t\t--> Outputs the Reverse Complement of the DNA of an input .fa file.")
        fmt.Println("\t binary-helix boyer-moore FILE PATTERN\t\t\t--> Performs Boyer-Moors searching algorithm on an input .fastq file")
        fmt.Println("\t binary-helix server\t\t\t\t\t--> Starts the server on port `4040`, turning your device into a supercomputer node.")
        fmt.Println("\t binary-helix admin_clear-db\t\t\t\t--> Clear EVERY item on the database. USE WITH CAUTION!")
        fmt.Println("")
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

	/* Start server */
	if strings.ToLower(arg[1]) == "server" {
		server.Server()
	}


	/* Admin_ Clear Database */
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
	// if strings.ToLower(arg[1]) == utils.AnalyserList[0] {
	if strings.ToLower(arg[1]) == "boyer-moore" {
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

	/* Complement */
	if strings.ToLower(arg[1]) == "complement" {
        if len(arg) > 3 {
			fmt.Println("FAIL: Please provide proper list of arguments!")
			fmt.Println("Type `binary-helix help` for more information.")

            return
        }

		path := arg[2]
        p := ""

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[1], p)

		// fmt.Println(lineCount)
		// fmt.Println(path)
		// fmt.Println(p)

		// pBM := types.ConstructBM(p)
		// fmt.Println(pBM.Bad_Character_Rule(2, "T"))
		// fmt.Println(analyser.BoyerMoore(p, pBM, t))
	}
}
