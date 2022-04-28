package client

import (
	"fmt"
	"net/rpc"
	"sync"

	"github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/utils"
)

/*
NOTE:
If you are catching the recieving data with a custom type (structure),
then you must define the type here too. Its best to share the type to
both server and client.
*/

func CheckServers() {
	var client *rpc.Client
	var err error

	nodeCount := utils.Get_ActiveNodeCount()
	var nodes []schema.Nodes

	for i := 0; i <= nodeCount; i++ {
		nodes = utils.Get_ActiveNodes()
	}

	for i := 0; i < nodeCount; i++ {
		client, err = rpc.DialHTTP("tcp", nodes[i].TargetIP_Port)
		if err != nil {
			utils.HandleError(err)
		}

		isAlive := "Dead"

		var response string

		client.Call("API.ImAlive", isAlive, &response)

		fmt.Printf("Response from %v ~~> %v\n", nodes[i].TargetIP_Port, response)
	}
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
		fmt.Println("FAIL: Server not responding!")
		fmt.Println("HELP: Please make sure the IP & Port are correct and that the server is running.")
		fmt.Println("NOTE: Please make sure that the port is open.")
		return false
	}

	fmt.Printf("Response from %v ~~> %v", ip_port, response)

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
		fmt.Println("FAIL: An individual host cannot register as more then 1 node!")
		return false
	}

	/* Insert node in compute_nodes collection */
	_, err = computeNodes.InsertOne(ctx, node)
	utils.HandleError(err)

	return true
}

func TaskServer(i string, s string, cId string, aArt schema.Analysis, extra string) {
	nodes := utils.Get_ActiveNodes()

	var inList []string
	inList = append(inList, i, s, cId, extra)

	for i, iL := range inList {
		fmt.Println(i, ">> ", iL)
	}
	fmt.Println("->", inList)
	fmt.Println("->", len(inList))
	fmt.Println("Task->", aArt.Task)

	for _, n := range nodes {
		client, err := rpc.DialHTTP("tcp", n.TargetIP_Port)
		if err != nil {
			utils.HandleError(err)
		}

		if aArt.Task == utils.AnalyserList[0] {
			var response string
			// var response []int
			client.Call("API.CallBoyerMoore", inList, &response)
			fmt.Println("|>", response)
			Merger(cId, aArt, response)
		}

		if aArt.Task == utils.AnalyserList[1] {
			var response string
			client.Call("API.CallComplement", inList[1], &response)
			fmt.Println("|>", response)
		}

		if aArt.Task == utils.AnalyserList[2] {
			var response string
			client.Call("API.CallReverseComplement", inList[1], &response)
			fmt.Println("|>", response)
		}

		if aArt.Task == utils.AnalyserList[3] {
			var response string
			client.Call("API.CallExactMatch", inList, &response)
			fmt.Println("|>", response)
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

	}

}

func Merger(cId string, aArt schema.Analysis, split string) {
	client, ctx := utils.ConnectDb()
	defer client.Disconnect(ctx)

	slicesDb := client.Database("slices_db")
	slices := slicesDb.Collection("slices")

    aArt.MergedOutput = split

	slice := schema.Slices{
		ComputationId: cId,
		SplitOrder:    1,
		Content:       split,
		AnalysisArt:   aArt,
	}

	/* Insert slice in slices collection */
	_, err := slices.InsertOne(ctx, slice)
	utils.HandleError(err)
}
