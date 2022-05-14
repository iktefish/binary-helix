package main

import (
	"fmt"
	"os"

	// "strconv"
	"strings"
	"sync"

	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/ui"
	// "github.com/iktefish/binary-helix/client"

	// "github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/server"
	"github.com/iktefish/binary-helix/types"
	"github.com/iktefish/binary-helix/utils"
	"github.com/iktefish/binary-helix/workers"
)

func main() {
	arg := os.Args

	// if len(arg) == 1 {
	// 	help()
	// 	return
	// }

	ui.Arg_Checker(arg)

	// if arg[1] == "RegisterNode" {
	// 	client.RegisterNode("172.17.0.3:4042", "binary-helix_c2")
	// }

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

	if arg[1] == "Read" {
		// path := "test/input/phix.fa"
		// path := "./test/input/sra_data.fastq"

		// path := "./test/input/small_sra_data.fastq"

		// fileExt, processed, lineCount := workers.Reader(path)
		// splits := workers.Splitter(fileExt, processed, lineCount)
		// workers.Carrier(splits, "quality-scores")
	}

	if arg[1] == "Qual" {
		path := "./sample/small_sra_data.fastq"

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		analyser.Id_SeqQual(splits)
	}

	if arg[1] == "DB" {
		// utils.Admin_EchoDbContents("nodes_db")

		itemCount := 0
		mergedItemCount := 0
		utils.Admin_EchoDbContents("slices_db", &itemCount, &mergedItemCount)

		fmt.Println("\t itemCount:\t\t", itemCount)
		fmt.Println("\t mergedItemCount:\t", mergedItemCount)

		// itemCount := utils.Admin_EchoDbContents("slices_db")
		// fmt.Println(itemCount)

		// utils.Admin_EchoDbContents("bench_db")
		// utils.Admin_EchoDbContents("Hello_DB")

		// utils.Admin_DummyInComputeNodes()
		utils.Admin_DummyInSlices()
		// utils.Admin_DummyInBenchmarks()

		// utils.Admin_EchoDbs()

		// utils.Admin_ClearDbAll("nodes_db")
		// utils.Admin_ClearDbAll("slices_db")
		// utils.Admin_ClearDbAll("bench_db")
		// utils.Admin_ClearDbAll("Hello")
		// utils.Admin_ClearDbAll("all")

		// schema.Test_TimeToPrim()
	}

	// /* Payment Resolver */ // NOTE: INCOMPLETE
	// if arg[1] == "Resolve" {
	//        workers.PayResolver()
	// }

	// HERE_IT_STARTS:
	/* Command line args */

	/* Help */
	if strings.ToLower(arg[1]) == "help" {
		help()
	}

	/* Register Node */
	// if strings.ToLower(arg[1]) == "register-node" {
	// 	ip_port := arg[2]
	// 	node_name := arg[3]
	//
	// 	out := false
	// 	if node_name != "" {
	// 		out = client.RegisterNode(ip_port, node_name)
	// 	}
	//
	// 	if out != true {
	// 		fmt.Println("FAIL:\t Registration failed!")
	// 	}
	// }

	// /* Check Nodes */
	// if strings.ToLower(arg[1]) == "check-nodes" {
	// 	client.CheckServers()
	// }

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
			fmt.Println("FAIL:\t Please provide which database you want purged!")
			fmt.Println("\t Possible options are:")
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

	/* Boyer-Moore (Node) */
	if strings.ToLower(arg[1]) == "boyer-moore-node" {
		if len(arg) != 4 {
			fmt.Println("FAIL:\t Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		p := arg[3]

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[0], p)
	}

	/* Boyer-Moore */
	// if strings.ToLower(arg[1]) == "boyer-moore" {
	// 	if len(arg) != 4 {
	// 		fmt.Println("FAIL:\t Please provide proper list of arguments!")
	// 		fmt.Println("\t Type `binary-helix help` for more information.")
	//
	// 		return
	// 	}
	//
	// 	path := arg[2]
	// 	p := arg[3]
	//
	// 	_, processed, _ := workers.Reader(path)
	// 	// splits := workers.Splitter(fileExt, processed, lineCount)
	// 	// workers.Carrier(splits, utils.AnalyserList[0], p)
	//
	// 	pBM := types.ConstructBM(p)
	// 	out := analyser.BoyerMoore(p, pBM, string(processed))
	// 	fmt.Println("OUTPUT:\t")
	// 	fmt.Println("\t ", out)
	// }

	/* Complement (Node) */
	if strings.ToLower(arg[1]) == "complement-node" {
		if len(arg) > 3 {
			fmt.Println("FAIL:\t  Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		p := ""

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[1], p)
	}

	/* Complement */
	if strings.ToLower(arg[1]) == "complement" {
		if len(arg) > 3 {
			fmt.Println("FAIL:\t  Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		// p := ""

		_, processed, _ := workers.Reader(path)
		// splits := workers.Splitter(fileExt, processed, lineCount)
		// workers.Carrier(splits, utils.AnalyserList[0], p)

		out := analyser.Complement(string(processed))
		fmt.Println("OUTPUT:\t")
		fmt.Println("\t ", out)
	}

	/* Reverse Complement (Node) */
	if strings.ToLower(arg[1]) == "reverse-complement-node" {
		if len(arg) > 3 {
			fmt.Println("FAIL:\t Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		p := ""

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[2], p)
	}

	/* Reverse Complement */
	if strings.ToLower(arg[1]) == "reverse-complement" {
		if len(arg) > 3 {
			fmt.Println("FAIL:\t Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		// p := ""

		_, processed, _ := workers.Reader(path)
		out := analyser.ReverseComplement(string(processed))
		fmt.Println("OUTPUT:\t")
		fmt.Println("\t ", out)
	}

	/* Exact Match (Node) */
	if strings.ToLower(arg[1]) == "exact-match-node" {
		if len(arg) != 4 {
			fmt.Println("FAIL:\t Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		p := arg[3]

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[3], p)
	}

	/* Exact Match */
	if strings.ToLower(arg[1]) == "exact-match" {
		if len(arg) != 4 {
			fmt.Println("FAIL:\t Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		p := arg[3]

		_, processed, _ := workers.Reader(path)
		out := analyser.ExactMatch(p, string(processed))
		fmt.Println("OUTPUT:\t")
		fmt.Println("\t ", out)
	}

	/* Longest Common Prefix */
	// if strings.ToLower(arg[1]) == "longest-common-prefix" {
	// 	if len(arg) != 4 {
	// 		fmt.Println("FAIL:\t Please provide proper list of arguments!")
	// 		fmt.Println("\t Type `binary-helix help` for more information.")

	// 		return
	// 	}

	// 	path_1 := arg[2]
	// 	path_2 := arg[3]

	// 	_, processed_1, _ := workers.Reader(path_1)
	// 	_, processed_2, _ := workers.Reader(path_2)

	// 	str_1 := string(processed_1)
	// 	str_2 := string(processed_2)

	// 	if len(str_1) != len(str_2) {
	// 		fmt.Println("FAIL:\t The DNA sequences are not from the same Genome!")
	// 		fmt.Println("\t The size of the DNA sequence must be equal to perform this match.")

	// 		return
	// 	}

	// 	lcp := analyser.LongestCommonPrefix(str_1, str_2)

	// 	fmt.Println(lcp)
	// }

	/* Total Number of Each Base */
	if strings.ToLower(arg[1]) == "total-bases-of-each" {
		if len(arg) != 3 {
			fmt.Println("FAIL:\t Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]

		As, Cs, Gs, Ts := analyser.TotalBasesOfEach(path)

		fmt.Println("\nOUTPUT:")
		fmt.Println("\t Number of Adenine(s): ", As)
		fmt.Println("\t Number of Cytosine(s): ", Cs)
		fmt.Println("\t Number of Guanine(s): ", Gs)
		fmt.Println("\t Number of Thymine(s): ", Ts)
		fmt.Println("")
	}

	/* K-Mer Index */
	if strings.ToLower(arg[1]) == "k-mer" {
		if len(arg) != 5 {
			fmt.Println("FAIL:\t Please provide proper list of arguments!")
			fmt.Println("\t Type `binary-helix help` for more information.")

			return
		}

		path := arg[2]
		from, err_1 := utils.Conv_StrToInt(arg[3])
		to, err_2 := utils.Conv_StrToInt(arg[4])

		if err_1 != nil || err_2 != nil {
			fmt.Println("FAIL:\t The last 2 arguments must be integers.")
			fmt.Println("\t They note `from` and `to` respectively, of the look-up-table you want to explore.")

			return
		}

		fileExt, processed, lineCount := workers.Reader(path)
		splits := workers.Splitter(fileExt, processed, lineCount)

		var wg sync.WaitGroup
		var mutex sync.Mutex

		isA := make([]analyser.IndexArt, len(splits))

		for i, s := range splits {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ia := analyser.ConstructIA(s, 2)

				mutex.Lock()
				fmt.Println("\nK-MER INDEX for split:", i+1)
				fmt.Println(ia.I[from:to])
				isA = append(isA, ia)
				mutex.Unlock()
			}()
			wg.Wait()
		}
	}

}

func help() {
	// fmt.Println("")
	// fmt.Println("INTRODUCTION:")
	// fmt.Println("")
	// fmt.Println("\t Welcome to Binary Helix, a distributed Genome analysis system powered by SliceCoin.")
	// fmt.Println("")
	// fmt.Println("\t Using this system, you can allow scientists, researchers and engineers from all over")
	// fmt.Println("\t the world to use a tiny fraction of you smartphone/desktop/leptops' computational power")
	// fmt.Println("\t for the purpose of analysing DNA sequences. Doing so you can be a part of may glorious")
	// fmt.Println("\t individuals who contribute to seeking, and helping others seek, a greater understanding of")
	// fmt.Println("\t things such as Cancer, Down's Syndrome, Aging, Genetic Psychiatric Conditions, Evolution,")
	// fmt.Println("\t Language, etc.")
	// fmt.Println("")
	// fmt.Println("\t You will of-course be paid for donating your computation to the service of science.")
	// fmt.Println("")
	// fmt.Println("COMMANDS:")
	// fmt.Println("")
	// fmt.Println("\t binary-helix help\t\t\t\t\t--> Outputs information on how to use this CLI")
	// fmt.Println("\t binary-helix register-node IP:PORT\t\t\t--> Registers your device as a server-node to donate computation.")
	// fmt.Println("\t binary-helix check-nodes\t\t\t\t--> Outputs the Complement of the DNA of an input .fa file.")
	// fmt.Println("\t binary-helix boyer-moore FILE PATTERN\t\t\t--> Performs Boyer-Moors searching algorithm on an input .fastq file")
	// fmt.Println("\t binary-helix boyer-moore-node FILE PATTERN\t\t--> Performs Boyer-Moors searching algorithm on an input .fastq file utilizing cluster.")
	// fmt.Println("\t binary-helix complement FILE\t\t\t\t--> Outputs the Complement of the DNA of an input .fa file.")
	// fmt.Println("\t binary-helix complement-node FILE\t\t\t--> Outputs the Complement of the DNA of an input .fa file utilizing cluster.")
	// fmt.Println("\t binary-helix reverse-complement FILE\t\t\t--> Outputs the Reverse Complement of the DNA of an input .fa file.")
	// fmt.Println("\t binary-helix reverse-complement-node FILE\t\t--> Outputs the Reverse Complement of the DNA of an input .fa file utilizing cluster.")
	// fmt.Println("\t binary-helix exact-match FILE PATTERN\t\t\t--> Outputs the Exact Match of the DNA read from an input .fa file.")
	// fmt.Println("\t binary-helix exact-match-node FILE PATTERN\t\t--> Outputs the Exact Match of the DNA read from an input .fa file utilizing cluster.")
	// fmt.Println("\t binary-helix k-mer FILE NUMBER\t\t\t\t--> Outputs the K-Mer Index of the DNA read from an input .fa file.")
	// fmt.Println("\t binary-helix longest-common-prefix FILE PATTERN\t--> Outputs the Longest Common Prefix the DNA read has with an input .fa file.")
	// fmt.Println("\t binary-helix total-bases-of-each FILE \t\t\t--> Outputs the Total Number of Each Base in the DNA read has with an input .fa file.")
	// fmt.Println("\t binary-helix server\t\t\t\t\t--> Starts the server on port `4040`, turning your device into a supercomputer node.")
	// fmt.Println("\t binary-helix admin_clear-db\t\t\t\t--> Clear EVERY item on the database. USE WITH CAUTION!")
	// fmt.Println("")
}
