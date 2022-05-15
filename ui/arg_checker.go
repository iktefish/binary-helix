package ui

import (
	"fmt"
	"strings"

	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/client"
	"github.com/iktefish/binary-helix/types"
	"github.com/iktefish/binary-helix/utils"
	"github.com/iktefish/binary-helix/workers"
)

// List of arguments expected by CLI.
var Valid_Args []string = []string{
	"",
	"help",
	"register-node",
	"check-nodes",
	"boyer-moore",
	"boyer-moore-node",
	"complement",
	"complement-node",
	"reverse-complement",
	"reverse-complement-node",
	"exact-match",
	"exact-match-node",
	"k-mer",
	"longest-common-prefix",
	"serve",
	"admin-clear-db",
	"blockchain-state",
	"restore-broken-computation",
}

// Take in the list of arguments passed by user and execute helper
// functions. Based on return values of helper functions, arg_checker() will either
// continue execution or panic.
func Arg_Checker(arg []string) {
	// Empty arguments
	if len(arg) == 1 {
		helper_empty()
		return
	}

	// Arbitrary argument
	if is_arg_present(arg) != true {
		fmt.Printf("\nFAIL:\t Invalid argument! '%v' is not a valid argument.\n", arg[1])
		fmt.Println("\t Type `binary-helix help` for more information.")
		fmt.Println()
		return
	}

	// Help
	if strings.ToLower(arg[1]) == Valid_Args[1] {
		if helper_help(arg) != true {
			fmt.Printf("\nFAIL:\t Invalid argument length! The `help` command should not have any trailing arguments.\n")
			fmt.Println("\t Type `binary-helix help` for more information.")
			fmt.Println()
			return
		}
		return
	}

	// Register node
	if strings.ToLower(arg[1]) == Valid_Args[2] {
		if helper_register_node(arg) != true {
			fmt.Printf("\nFAIL:\t Invalid argument length! The `register-node` command expects 'IP_Address:Port' and 'Node name'.\n")
			fmt.Println("\t Type `binary-helix help` for more information.")
			fmt.Println()
			return
		}
		return
	}

	// Check nodes
	if strings.ToLower(arg[1]) == Valid_Args[3] {
		if helper_check_nodes(arg) != true {
			fmt.Printf("\nFAIL:\t Invalid argument length! The `check-nodes` command should not have any trailing arguments.\n")
			fmt.Println("\t Type `binary-helix help` for more information.")
			fmt.Println()
			return
		}
		return
	}

	// Boyer Moore
	if strings.ToLower(arg[1]) == Valid_Args[4] {
		catch_bool, catch_code := helper_boyer_moore(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `boyer-moore` command expects 2 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				fmt.Printf("\t\t 2. A pattern (string) to match within the reference genome.\n")
				fmt.Println()
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `boyer-moore` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			case 2:
				fmt.Printf("\nFAIL:\t Invalid argument! Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			case 3:
				fmt.Printf("\nFAIL:\t Invalid argument!\n")
				fmt.Println()
				fmt.Printf("\t Patterns cannot contain numbers or symbols.\n")
				fmt.Printf("\t Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			}
			return
		}
		return
	}

	// Boyer Moore (node)
	if strings.ToLower(arg[1]) == Valid_Args[5] {
		catch_bool, catch_code := helper_boyer_moore_node(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `boyer-moore-node` command expects 2 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				fmt.Printf("\t\t 2. A pattern (string) to match within the reference genome.\n")
				fmt.Println()
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `boyer-moore` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			case 2:
				fmt.Printf("\nFAIL:\t Invalid argument! Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			case 3:
				fmt.Printf("\nFAIL:\t Invalid argument!\n")
				fmt.Println()
				fmt.Printf("\t Patterns cannot contain numbers or symbols.\n")
				fmt.Printf("\t Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			}
			return
		}
		return
	}

	// Complement
	if strings.ToLower(arg[1]) == Valid_Args[6] {
		catch_bool, catch_code := helper_complement(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `complement` command expects 1 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `complement` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			}
		}
	}

	// Complement (node)
	if strings.ToLower(arg[1]) == Valid_Args[7] {
		catch_bool, catch_code := helper_complement_node(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `complement` command expects 1 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `complement` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			}
		}
	}

	// Reverse Complement
	if strings.ToLower(arg[1]) == Valid_Args[8] {
		catch_bool, catch_code := helper_reverse_complement(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `reverse-complement` command expects 1 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `reverse-complement` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			}
		}
	}

	// Reverse Complement (node)
	if strings.ToLower(arg[1]) == Valid_Args[9] {
		catch_bool, catch_code := helper_reverse_complement_node(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `reverse-complement-node` command expects 1 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `reverse-complement-node` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			}
		}
	}

	// (Naive) Exact Match
	if strings.ToLower(arg[1]) == Valid_Args[10] {
		catch_bool, catch_code := helper_exact_match(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `exact-match` command expects 2 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				fmt.Printf("\t\t 2. A pattern (string) to match within the reference genome.\n")
				fmt.Println()
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `exact-match` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			case 2:
				fmt.Printf("\nFAIL:\t Invalid argument! Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			case 3:
				fmt.Printf("\nFAIL:\t Invalid argument!\n")
				fmt.Println()
				fmt.Printf("\t Patterns cannot contain numbers or symbols.\n")
				fmt.Printf("\t Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			}
			return
		}
		return
	}

	// (Naive) Exact Match (node)
	if strings.ToLower(arg[1]) == Valid_Args[11] {
		catch_bool, catch_code := helper_exact_match_node(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `exact-match-node` command expects 2 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				fmt.Printf("\t\t 2. A pattern (string) to match within the reference genome.\n")
				fmt.Println()
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `exact-match-node` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			case 2:
				fmt.Printf("\nFAIL:\t Invalid argument! Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			case 3:
				fmt.Printf("\nFAIL:\t Invalid argument!\n")
				fmt.Println()
				fmt.Printf("\t Patterns cannot contain numbers or symbols.\n")
				fmt.Printf("\t Patterns can only contain the following characters (regardless of case):\n")
				fmt.Println()
				fmt.Printf("\t\t 1. 'A'\n")
				fmt.Printf("\t\t 2. 'C'\n")
				fmt.Printf("\t\t 3. 'G'\n")
				fmt.Printf("\t\t 4. 'T'\n")
				fmt.Println()
				return
			}
			return
		}
		return
	}

	// K-Mer Index
	if strings.ToLower(arg[1]) == Valid_Args[12] {
		catch_bool, catch_code := helper_k_mer(arg)
		if catch_bool != true {
			switch catch_code {
			case 0:
				fmt.Printf("\nFAIL:\t Invalid argument length! The `k-mer` command expects 3 arguments:\n")
				fmt.Println()
				fmt.Printf("\t\t 1. Path to a `.fa` file containing reference genome.\n")
				fmt.Printf("\t\t 2. A number denoting 'FROM'.\n")
				fmt.Printf("\t\t 2. A number denoting 'TO'.\n")
				fmt.Println()
				return
			case 1:
				fmt.Printf("\nFAIL:\t Invalid argument! First argument of the `exact-match-node` command must be path to a `.fa` file.\n")
				fmt.Println()
				return
			case 2:
				fmt.Println("\nFAIL:\t Invalid arguments. The last 2 arguments must be integers.")
				fmt.Println("\t They note FROM and TO respectively, of the look-up-table you want to explore.")
				fmt.Println()
				return
			case 3:
				fmt.Println("\nFAIL:\t Invalid FROM and/or TO arguments.")
				fmt.Println("\t Argument FROM cannot be greater than or equal to argument TO.")
				fmt.Println()
				return
			}
			return
		}
		return
	}
}

// Executing binary with no arguments is an alias to `binary-helix help`.
func helper_empty() bool {
	helper_help([]string{"binary-helix", "help"})
	return true
}

// Shows help menu. Execute when `binary-helix help` is invoked.
func helper_help(arg []string) bool {
	if len(arg) != 2 {
		return false
	}

	fmt.Println("")
	fmt.Println("INTRODUCTION:")
	fmt.Println("")
	fmt.Println("\t Welcome to Binary Helix, a distributed Genome analysis system powered by SliceCoin.")
	fmt.Println("")
	fmt.Println("\t Using this system, you can allow scientists, researchers and engineers from all over")
	fmt.Println("\t the world to use a tiny fraction of you smartphones'/desktops'/laptops' computational power")
	fmt.Println("\t for the purpose of analysing DNA sequences. Doing so you can be a part of may glorious")
	fmt.Println("\t individuals who contribute to seeking, and helping others seek, a greater understanding of")
	fmt.Println("\t things such as cancer, Down's syndrome, aging, genetic and psychiatric conditions, evolution,")
	fmt.Println("\t Language, etc.")
	fmt.Println("")
	fmt.Println("\t You will of-course be paid for donating your computation to the service of science.")
	fmt.Println("")
	fmt.Println("COMMANDS:")
	fmt.Println("")
	fmt.Println("\t binary-helix help\t\t\t\t\t--> Outputs information on how to use this CLI")
	fmt.Println("\t binary-helix register-node IP:PORT NAME\t\t--> Registers your device as a server-node to donate computation.")
	fmt.Println("\t binary-helix check-nodes\t\t\t\t--> Outputs the Complement of the DNA of an input .fa file.")
	fmt.Println("\t binary-helix boyer-moore FILE PATTERN\t\t\t--> Performs Boyer-Moors searching algorithm on an input .fastq file")
	fmt.Println("\t binary-helix boyer-moore-node FILE PATTERN\t\t--> Performs Boyer-Moors searching algorithm on an input .fastq file utilizing cluster.")
	fmt.Println("\t binary-helix complement FILE\t\t\t\t--> Outputs the Complement of the DNA of an input .fa file.")
	fmt.Println("\t binary-helix complement-node FILE\t\t\t--> Outputs the Complement of the DNA of an input .fa file utilizing cluster.")
	fmt.Println("\t binary-helix reverse-complement FILE\t\t\t--> Outputs the Reverse Complement of the DNA of an input .fa file.")
	fmt.Println("\t binary-helix reverse-complement-node FILE\t\t--> Outputs the Reverse Complement of the DNA of an input .fa file utilizing cluster.")
	fmt.Println("\t binary-helix exact-match FILE PATTERN\t\t\t--> Outputs the Exact Match of the DNA read from an input .fa file.")
	fmt.Println("\t binary-helix exact-match-node FILE PATTERN\t\t--> Outputs the Exact Match of the DNA read from an input .fa file utilizing cluster.")
	fmt.Println("\t binary-helix k-mer FILE TO FROM\t\t\t--> Outputs the K-Mer Index of the DNA read from an input .fa file.")
	fmt.Println("\t binary-helix longest-common-prefix FILE PATTERN\t--> Outputs the Longest Common Prefix the DNA read has with an input .fa file.")
	fmt.Println("\t binary-helix total-bases-of-each FILE \t\t\t--> Outputs the Total Number of Each Base in the DNA read has with an input .fa file.")
	fmt.Println("\t binary-helix server\t\t\t\t\t--> Starts the server on port `4040`, turning your device into a supercomputer node.")
	fmt.Println("\t binary-helix admin_clear-db\t\t\t\t--> Clear EVERY item on the database. USE WITH CAUTION!")
	fmt.Println("")

	return true
}

// Allows the master node to register a server by providing the IP address of
// said node.
func helper_register_node(arg []string) bool {
	if len(arg) != 4 {
		return false
	}

	ip_port := arg[2]
	node_name := arg[3]
	out := false

	if node_name == "" {
		return false
	}

	out = client.RegisterNode(ip_port, node_name)

	if out != true {
		fmt.Println("FAIL:\t Registration failed!")
	}

	return true
}

// Ping and check status of all registered node.
func helper_check_nodes(arg []string) bool {
	if len(arg) != 2 {
		return false
	}

	client.Check_Server()
	return true
}

// Performs Boyer Moore exact pattern match on an input genome.
// Make sure input file is a proper `.fa` file format. Example
// pattern would be "TCCA".
func helper_boyer_moore(arg []string) (bool, int) {
	if len(arg) != 4 {
		return false, 0
	}

	path := arg[2]
	pattern := strings.ToUpper(arg[3])

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	catch_bool, catch_code := is_pattern_correct(pattern)
	if catch_bool != true && catch_code == 0 {
		return false, 2
	}
	if catch_bool != true && catch_code == 1 {
		return false, 3
	}

	_, processed, _ := workers.Reader(path)

	pBM := types.ConstructBM(pattern)
	out := analyser.BoyerMoore(pattern, pBM, string(processed))

	fmt.Println("\nOUTPUT:\t")
	fmt.Println()
	fmt.Print("\t")
	for _, output := range out {
		fmt.Print(" ", output)
	}
	fmt.Print("\n\n")
	return true, 0
}

// Performs Boyer Moore exact pattern match using available nodes on an
// input genome. Make sure input file is a proper `.fa` file format. Example
// pattern would be "TCCA".
func helper_boyer_moore_node(arg []string) (bool, int) {
	if len(arg) != 4 {
		return false, 0
	}

	path := arg[2]
	pattern := strings.ToUpper(arg[3])

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	catch_bool, catch_code := is_pattern_correct(pattern)
	if catch_bool != true && catch_code == 0 {
		return false, 2
	}
	if catch_bool != true && catch_code == 1 {
		return false, 3
	}

	file_ext, processed, line_count := workers.Reader(path)
	splits := workers.Splitter(file_ext, processed, line_count)
	workers.Carrier(splits, utils.AnalyserList[0], pattern)

	/*
		fmt.Println("\nOUTPUT:\t")
		fmt.Println()
		fmt.Print("\t")
		for _, output := range out {
			fmt.Print(" ", output)
		}
		fmt.Print("\n\n")
	*/
	return true, 0
}

// Computes complement strand of an input genome. Make sure input file is a
// proper `.fa` file.
func helper_complement(arg []string) (bool, int) {
	if len(arg) != 3 {
		return false, 0
	}

	path := arg[2]

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	_, processed, _ := workers.Reader(path)
	// splits := workers.Splitter(fileExt, processed, lineCount)
	// workers.Carrier(splits, utils.AnalyserList[1], "")

	out := analyser.Complement(string(processed))
	fmt.Println("\nOUTPUT:\t")
	fmt.Println("\t ", out)
	fmt.Println()
	return true, 0
}

// Computes complement strand using available nodes of an input genome.
// Make sure input file is a proper `.fa` file.
func helper_complement_node(arg []string) (bool, int) {
	if len(arg) != 3 {
		return false, 0
	}

	path := arg[2]

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	file_ext, processed, line_count := workers.Reader(path)
	splits := workers.Splitter(file_ext, processed, line_count)
	workers.Carrier(splits, utils.AnalyserList[1], "")

	// out := analyser.Complement(string(processed))
	// fmt.Println("\nOUTPUT:\t")
	// fmt.Println("\t ", out)
	// fmt.Println()
	return true, 0
}

// Computes reverse complement strand of an input genome. Make sure input file
// is a proper `.fa` file.
func helper_reverse_complement(arg []string) (bool, int) {
	if len(arg) != 3 {
		return false, 0
	}

	path := arg[2]

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	_, processed, _ := workers.Reader(path)
	// splits := workers.Splitter(fileExt, processed, lineCount)
	// workers.Carrier(splits, utils.AnalyserList[2], "")

	out := analyser.ReverseComplement(string(processed))
	fmt.Println("\nOUTPUT:\t")
	fmt.Println("\t ", out)
	fmt.Println()
	return true, 0
}

// Computes reverse complement strand using available nodes of an input genome.
// Make sure input file is a proper `.fa` file.
func helper_reverse_complement_node(arg []string) (bool, int) {
	if len(arg) != 3 {
		return false, 0
	}

	path := arg[2]

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	file_ext, processed, line_count := workers.Reader(path)
	splits := workers.Splitter(file_ext, processed, line_count)
	workers.Carrier(splits, utils.AnalyserList[2], "")

	// out := analyser.ReverseComplement(string(processed))
	// fmt.Println("\nOUTPUT:\t")
	// fmt.Println("\t ", out)
	// fmt.Println()
	return true, 0
}

// Computes (naive) exact matches of a pattern in an input genome. Make sure input file
// is a proper `.fa` file. Example pattern would be "TCCA".
func helper_exact_match(arg []string) (bool, int) {
	if len(arg) != 4 {
		return false, 0
	}

	path := arg[2]
	pattern := strings.ToUpper(arg[3])

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	catch_bool, catch_code := is_pattern_correct(pattern)
	if catch_bool != true && catch_code == 0 {
		return false, 2
	}
	if catch_bool != true && catch_code == 1 {
		return false, 3
	}

	_, processed, _ := workers.Reader(path)
	out := analyser.ExactMatch(pattern, string(processed))

	fmt.Println("\nOUTPUT:\t")
	fmt.Println()
	fmt.Print("\t")
	for _, output := range out {
		fmt.Print(" ", output)
	}
	fmt.Print("\n\n")
	return true, 0
}

// Computes (naive) exact matches of a pattern using available nodes in an input genome.
// Make sure input file is a proper `.fa` file. Example pattern would be "TCCA".
func helper_exact_match_node(arg []string) (bool, int) {
	if len(arg) != 4 {
		return false, 0
	}

	path := arg[2]
	pattern := strings.ToUpper(arg[3])

	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	catch_bool, catch_code := is_pattern_correct(pattern)
	if catch_bool != true && catch_code == 0 {
		return false, 2
	}
	if catch_bool != true && catch_code == 1 {
		return false, 3
	}

	file_ext, processed, line_count := workers.Reader(path)
	splits := workers.Splitter(file_ext, processed, line_count)
	workers.Carrier(splits, utils.AnalyserList[3], pattern)

	/*
		fmt.Println("\nOUTPUT:\t")
		fmt.Println()
		fmt.Print("\t")
		for _, output := range out {
			fmt.Print(" ", output)
		}
		fmt.Print("\n\n")
	*/
	return true, 0
}

// Computes k-mer index lookup table of input genome. Make sure input file is a
// proper `.fa` file. Will have a `--view` flag to simply just view the output
// table instead of storing it.
func helper_k_mer(arg []string) (bool, int) {
	if len(arg) != 5 {
		return false, 0
	}

	path := arg[2]
	if utils.Verify_Fasta(path) != true {
		return false, 1
	}

	from, err_1 := utils.Conv_StrToInt(arg[3])
	to, err_2 := utils.Conv_StrToInt(arg[4])
	if err_1 != nil || err_2 != nil {
		return false, 2
	}

	if from >= to {
		return false, 3
	}

	file_ext, processed, line_count := workers.Reader(path)
	splits := workers.Splitter(file_ext, processed, line_count)

	isA := make([]analyser.IndexArt, len(splits))
	for i, s := range splits {
		ia := analyser.ConstructIA(s, 2)
		fmt.Printf("\nK-MER INDEX [ %v ]\n", i+1)
		fmt.Println()
		fmt.Print("\t")
		for _, output := range ia.I[from:to] {
			fmt.Print(" ", output)
		}
		fmt.Println()
		isA = append(isA, ia)
	}
	fmt.Println()
	return true, 0
}

// Computes the longest common prefix shared between 2 genome of the same species.
// Make sure the input file is a proper `.fa` file.
func helper_longest_common_prefix() bool {
	return false
}

// Determines the total count of each bases present in a genome. Make sure the input
// file is a proper `.fa` file.
func helper_total_bases_of_each() bool {
	return false
}

// Execute server.
func helper_server() bool {
	return false
}

// Echo current blockchain state for SliceCoin in std::out.
func helper_blockchain_state() bool {
	return false
}

func helper_restore_broken_state() bool {
	return false
}

// WARNING: For admin/debug purposes
//
// Clear every documents in all databases.
func helper_admin_clear_db() bool {
	return false
}

// Check if input arg contains an expected arg.
func is_arg_present(arg []string) bool {
	for _, v := range Valid_Args {
		for _, a := range arg {
			if a == v {
				return true
			}
		}
	}

	return false
}

// Verify if input pattern only contains "ATCG". Returns:
//
// `false, 0` if pattern contains characters other than "ACGT",
// `false, 1` if pattern contains numerics or symbols,
// `true, 0` if pattern is correct.
func is_pattern_correct(pattern string) (bool, int) {
	pattern = strings.ToLower(pattern)
	if strings.Contains(pattern, "a") || strings.Contains(pattern, "t") || strings.Contains(pattern, "c") || strings.Contains(pattern, "g") {
		for _, char := range pattern {
			if (char > 'a' && char < 'c') || (char > 'c' && char < 'g') || (char > 'g' && char < 't') || (char > 't') {
				return false, 0
			}
			if char < 'a' || char > 'z' {
				return false, 1
			}
		}
		return true, 0
	}
	return false, 0
}
