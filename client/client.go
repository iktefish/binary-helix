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

// func Client() {
// 	client, err := rpc.DialHTTP("tcp", "172.17.0.2:4040")
// 	if err != nil {
// 		log.Fatal("Connection error: ", err)
// 	}
//
// 	s := "String from the client!"
// 	var response string
//
// 	client.Call("API.GetByName", s, &response)
// fmt.Println("response ~~> ", response) }

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
        fmt.Println(i,">> ", iL)
    }
	fmt.Println("->", inList)
	fmt.Println("->", len(inList))

	for _, n := range nodes {
		client, err := rpc.DialHTTP("tcp", n.TargetIP_Port)
		if err != nil {
			utils.HandleError(err)
		}

		var response string

		if aArt.Task == utils.AnalyserList[0] {
			client.Call("API.CallBoyerMoore", inList, &response)
		}

		fmt.Println("[0] :::", response)

		if aArt.Task == utils.AnalyserList[1] {
			client.Call("API.CallComplement", inList[1], &response)
		}

		fmt.Println("[1] :::", response)
	}

}
