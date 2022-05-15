package main

import (
	// "fmt"
	"os"

	// "strings"

	// "github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/ui"
	// "github.com/iktefish/binary-helix/utils"
	// "github.com/iktefish/binary-helix/workers"
)

func main() {
	arg := os.Args

	ui.Arg_Checker(arg)

	// if arg[1] == "DB" {
	// 	/* utils.Admin_EchoDbContents("nodes_db") */
	//
	// 	itemCount := 0
	// 	mergedItemCount := 0
	// 	utils.Admin_EchoDbContents("slices_db", &itemCount, &mergedItemCount)
	//
	// 	fmt.Println("\t itemCount:\t\t", itemCount)
	// 	fmt.Println("\t mergedItemCount:\t", mergedItemCount)
	//
	// 	/* itemCount := utils.Admin_EchoDbContents("slices_db")
	// 	fmt.Println(itemCount)
	//
	// 	utils.Admin_EchoDbContents("bench_db")
	// 	utils.Admin_EchoDbContents("Hello_DB")
	//
	// 	utils.Admin_DummyInComputeNodes() */
	// 	utils.Admin_DummyInSlices()
	// 	/* utils.Admin_DummyInBenchmarks()
	//
	// 	utils.Admin_EchoDbs()
	//
	// 	utils.Admin_ClearDbAll("nodes_db")
	// 	utils.Admin_ClearDbAll("slices_db")
	// 	utils.Admin_ClearDbAll("bench_db")
	// 	utils.Admin_ClearDbAll("Hello")
	// 	utils.Admin_ClearDbAll("all")
	//
	// 	schema.Test_TimeToPrim() */
	// }
	//
	// /* // Payment Resolver
	//     // NOTE: INCOMPLETE
	// 	if arg[1] == "Resolve" {
	// 	       workers.PayResolver()
	// 	} */
	//
	// // Admin_ Clear Database
	// if strings.ToLower(arg[1]) == "admin_clear-db" {
	// 	dbs := [4]string{
	// 		"nodes_db",
	// 		"slices_db",
	// 		"bench_db",
	// 		"all",
	// 	}
	//
	// 	if len(arg) < 3 {
	// 		fmt.Println("FAIL:\t Please provide which database you want purged!")
	// 		fmt.Println("\t Possible options are:")
	// 		for i := range dbs {
	// 			fmt.Println(i+1, dbs[i])
	// 		}
	//
	// 		return
	// 	}
	//
	// 	db := arg[2]
	// 	if db == dbs[0] {
	// 		utils.Admin_ClearDbAll("nodes_db")
	// 	}
	// 	if db == dbs[1] {
	// 		utils.Admin_ClearDbAll("slices_db")
	// 	}
	// 	if db == dbs[2] {
	// 		utils.Admin_ClearDbAll("bench_db")
	// 	}
	// 	if db == dbs[3] {
	// 		utils.Admin_ClearDbAll("all")
	// 	}
	// }
}
