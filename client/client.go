package client

import (
	"fmt"
	"log"
	"net/rpc"

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
//
// 	fmt.Println("response ~~> ", response)
// }

func CheckServers() {
	var client *rpc.Client
	var err error

	nodeCount := utils.Get_ActiveNodeCount()
	var nodes []schema.Nodes

	for i := 0; i <= nodeCount; i++ {
		nodes = utils.Get_ActiveNodes()
	}

	for i := 0; i <= nodeCount; i++ {
		client, err = rpc.DialHTTP("tcp", nodes[i].TargetIP_Port)
		if err != nil {
			utils.HandleError(err)
		}

		isAlive := false
		var response bool

		client.Call("API.ImAlive", isAlive, &response)

		if response != true {
			fmt.Println("FAIL: Server not responding!")
			log.Fatal("FAIL: Node appears to be dead!")
		}

		fmt.Printf("Response from %v ~~> %v", nodes[i].TargetIP_Port, response)
	}
}

func RegisterNode(ip_port string, node_name string) bool {
	client, err := rpc.DialHTTP("tcp", ip_port)
	if err != nil {
		utils.HandleError(err)
	}

	isAlive := false
	var response bool

	client.Call("API.ImAlive", isAlive, &response)
	if response != true {
		fmt.Println("FAIL: Server not responding!")
		fmt.Println("HELP: Please make sure the IP:Port is correct and that the server is running.")
		fmt.Println("NOTE: Please make sure that the port is open.")
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

	/* Insert node in compute_nodes collection */
	_, err = computeNodes.InsertOne(ctx, node)
	utils.HandleError(err)

	return true
}
