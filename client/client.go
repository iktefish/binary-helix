package client

import (
	"fmt"
	"net/rpc"
	"sync"

	"github.com/iktefish/binary-helix/blockchain"
	"github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/utils"
)

/*
NOTE:
If you are catching the recieving data with a custom type (structure),
then you must define the type here too. Its best to share the type to
both server and client.
*/

// Query database to retrieve list of active servers, and ping each
// server to check their status.
func Check_Server() {
	var client *rpc.Client
	var err error

	nodeCount := utils.Get_ActiveNodeCount()
	if nodeCount == 0 {
		fmt.Println("\nOUTPUT:")
		fmt.Println("\t There are no registered node.")
		fmt.Println("\t To register a node, run `binary-helix register-node IP:PORT NODE_NAME`.")
		fmt.Println()
		return
	}

	var nodes []schema.Nodes

	for i := 0; i <= nodeCount; i++ {
		nodes = utils.Get_ActiveNodes()
	}

	fmt.Printf("\nResponse:\n")
	fmt.Println()
	for i := 0; i < nodeCount; i++ {
		client, err = rpc.DialHTTP("tcp", nodes[i].TargetIP_Port)
		if err != nil {
			fmt.Println("\nDANGER:")
			fmt.Println("\t One or more nodes are dead. Please troubleshoot.")
			fmt.Println()
			return
			// utils.HandleError(err)
		}

		isAlive := "Dead"
		var response string
		client.Call("API.ImAlive", isAlive, &response)

		fmt.Printf("\t • NAME: %v\t ADDRESS: %v \t~~> %v\n", nodes[i].NodeName, nodes[i].TargetIP_Port, response)
	}
	fmt.Println()
}

func RegisterNode(ip_port string, node_name string) bool {
	var wg sync.WaitGroup
	wg.Add(1)

	client, err := rpc.DialHTTP("tcp", ip_port)
	if err != nil {
		utils.HandleError(err)
	}

	isAlive := "Dead"
	var response string

	go func() {
		defer wg.Done()
		client.Call("API.ImAlive", isAlive, &response)
	}()
	wg.Wait()

	if response != "Alive" {
		fmt.Println("FAIL:\t Server not responding!")
		fmt.Println("HELP:\t Please make sure the IP & Port are correct and that the server is running.")
		fmt.Println("NOTE:\t Please make sure that the port is open.")
		return false
	}

	dbclient, ctx := utils.ConnectDb()
	defer dbclient.Disconnect(ctx)

	nodesDb := dbclient.Database("nodes_db")
	computeNodes := nodesDb.Collection("compute_nodes")

	node := schema.Nodes{
		NodeName:            node_name,
		TargetIP_Port:       ip_port,
		TotalCreditAttained: 0,
		Active:              true,
	}

	if utils.Verify_NodeNoDup(ip_port, node_name) != true {
		return false
	}

	/* Insert node in compute_nodes collection */
	_, err = computeNodes.InsertOne(ctx, node)
	utils.HandleError(err)

	return true
}

func TaskServer(s []string, cId string, aArt []schema.Analysis, extra string) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	nodes := utils.Get_ActiveNodes()

	var responses []string

	// var inList []string
	// inList = append(inList, i, s, cId, extra)

	// for i, iL := range inList {
	// 	fmt.Println(i, ">> ", iL)
	// }
	// fmt.Println("->", inList)
	// fmt.Println("->", len(inList))
	// fmt.Println("Task->", aArt.Task)

	for n, node := range nodes {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var inList []string
			inList = append(inList, "", s[n], cId, extra)

			client, err := rpc.DialHTTP("tcp", node.TargetIP_Port)
			if err != nil {
				utils.HandleError(err)
			}

			if aArt[n].Task == utils.AnalyserList[0] {
				var response string
				client.Call("API.CallBoyerMoore", inList, &response)
				// fmt.Println("|>", response)
				// fmt.Println(node.TargetIP_Port)
				// fmt.Println(nodes)
				Merger(cId, aArt[n], response)
				mu.Lock()
				responses = append(responses, response)

				chain := blockchain.InitBlockChain()
				chain.AddBlock(cId, aArt[n].TargetIP_Port)
				fmt.Printf("\n")
				fmt.Printf("\t %v:\n\n", node.NodeName)
				for i, block := range chain.Blocks {
					fmt.Printf("\t\tINDEX %v |--> \tblock.C_Id\t~~>\t%x\n", i, block.C_Id)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Ip_Port\t~~>\t%x\n", i, block.Ip_Port)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Hash\t~~>\t%x\n", i, block.Hash)
				}
				mu.Unlock()
			}

			if aArt[n].Task == utils.AnalyserList[1] {
				var response string
				client.Call("API.CallComplement", inList[1], &response)
				// fmt.Println("|>", response)
				// fmt.Println(node.TargetIP_Port)
				// fmt.Println(nodes)
				Merger(cId, aArt[n], response)
				mu.Lock()
				responses = append(responses, response)

				chain := blockchain.InitBlockChain()
				chain.AddBlock(cId, aArt[n].TargetIP_Port)
				fmt.Printf("\n")
				fmt.Printf("\t %v:\n\n", node.NodeName)
				for i, block := range chain.Blocks {
					fmt.Printf("\t\tINDEX %v |--> \tblock.C_Id\t~~>\t%x\n", i, block.C_Id)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Ip_Port\t~~>\t%x\n", i, block.Ip_Port)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Hash\t~~>\t%x\n", i, block.Hash)
				}
				mu.Unlock()
			}

			if aArt[n].Task == utils.AnalyserList[2] {
				var response string
				client.Call("API.CallReverseComplement", inList[1], &response)
				// fmt.Println("|>", response)
				// fmt.Println(node.TargetIP_Port)
				// fmt.Println(nodes)
				Merger(cId, aArt[n], response)
				mu.Lock()
				responses = append(responses, response)

				chain := blockchain.InitBlockChain()
				chain.AddBlock(cId, aArt[n].TargetIP_Port)
				fmt.Printf("\n")
				fmt.Printf("\t %v:\n\n", node.NodeName)
				for i, block := range chain.Blocks {
					fmt.Printf("\t\tINDEX %v |--> \tblock.C_Id\t~~>\t%x\n", i, block.C_Id)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Ip_Port\t~~>\t%x\n", i, block.Ip_Port)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Hash\t~~>\t%x\n", i, block.Hash)
				}
				mu.Unlock()
			}

			if aArt[n].Task == utils.AnalyserList[3] {
				var response string
				client.Call("API.CallExactMatch", inList, &response)
				// fmt.Println("|>", response)
				// fmt.Println(node.TargetIP_Port)
				// fmt.Println(nodes)
				Merger(cId, aArt[n], response)
				mu.Lock()
				responses = append(responses, response)

				chain := blockchain.InitBlockChain()
				chain.AddBlock(cId, aArt[n].TargetIP_Port)
				fmt.Printf("\n")
				fmt.Printf("\t %v:\n\n", node.NodeName)
				for i, block := range chain.Blocks {
					fmt.Printf("\t\tINDEX %v |--> \tblock.C_Id\t~~>\t%x\n", i, block.C_Id)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Ip_Port\t~~>\t%x\n", i, block.Ip_Port)
					fmt.Printf("\t\tINDEX %v |--> \tblock.Hash\t~~>\t%x\n", i, block.Hash)
				}
				mu.Unlock()
			}

			// /* INCOMPLETE */
			// if aArt.Task == utils.AnalyserList[4] {
			//  var response string
			// 	client.Call("API.CallKMerIndex", inList, &response)
			// }
			//
			// if aArt.Task == utils.AnalyserList[5] {
			//  var response string
			// 	client.Call("API.CallLongestCommonPrefix", inList, &response)
			// }
			//
			// if aArt.Task == utils.AnalyserList[7] {
			//  var response string
			// 	client.Call("API.CallTotalBasesOfEach", inList[3], &response)
			// }
			// if aArt.Task == utils.AnalyserList[8] {
			//  var response string
			// 	client.Call("API.CallIdQual", inList, &response)
			// }
		}()
		wg.Wait()
	}

	fmt.Println("\nOUTPUT:\t")
	fmt.Println()
	fmt.Print("\t")
	for _, response := range responses {
		if aArt[0].Task == utils.AnalyserList[0] || aArt[0].Task == utils.AnalyserList[3] {
			fmt.Print(" ", response)
		}
		if aArt[0].Task == utils.AnalyserList[1] || aArt[0].Task == utils.AnalyserList[2] {
			fmt.Print("", response)
		}
	}
	fmt.Print("\n\n")
}

func Merger(cId string, aArt schema.Analysis, split string) {
	client, ctx := utils.ConnectDb()
	defer client.Disconnect(ctx)

	slicesDb := client.Database("slices_db")
	slices := slicesDb.Collection("slices")

	// aArt.MergedOutput = split

	var i int32

	var responses []string
	responses = append(responses, split)

	slice := schema.Slices{
		ComputationId: cId,
		SplitOrder:    i,
		Content:       split,
		AnalysisArt:   aArt,
		MergedOutput:  responses,
	}

	/* Insert slice in slices collection */
	_, err := slices.InsertOne(ctx, slice)
	utils.HandleError(err)
}
