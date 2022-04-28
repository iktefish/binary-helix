package utils

import (
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckNodeDup() bool {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	/* Catch nodesDb database and compute_nodes collection */
	nodesDb := client.Database("nodes_db")
	computeNodes := nodesDb.Collection("compute_nodes")

	/* NOTE: This will be taken from user-input */
	RegisterNode := schema.Nodes{
		NodeName:                "binary-helix_c1",
		TargetIP_Port:                "172.17.0.2",
		Active:                  true,
		TotalCreditAttained:     0,
		TaskCompletionFrequency: 0,
	}

	/* Check for duplicate IP */
	cursor, err := computeNodes.Find(ctx, bson.M{"target_ip_port": RegisterNode.TargetIP_Port})
	HandleError(err)
	defer cursor.Close(ctx)
	var results []schema.Nodes
	if err = cursor.All(ctx, &results); err != nil {
		HandleError(err)
	}

	if len(results) > 0 {
		return true
	} else {
		return false
	}
}
