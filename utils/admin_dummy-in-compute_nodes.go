package utils

import (
	"fmt"
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func Admin_DummyInComputeNodes() {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	nodesDb := client.Database("nodes_db")
	computeNodes := nodesDb.Collection("compute_nodes")

	dummy := schema.Nodes{
		NodeName:                "DUMMY",
		TargetIP_Port:                "DUMMY",
		Active:                  true,
		TotalCreditAttained:     0,
		TaskCompletionFrequency: 0,
	}

	/* Insert dummy in computeNodes collection */
	nodeInsertResult, err := computeNodes.InsertOne(ctx, dummy)
	HandleError(err)

	/* Query computeNodes collection */
	cursor, err := computeNodes.Find(ctx, bson.M{"_id": nodeInsertResult.InsertedID})
	HandleError(err)

	defer cursor.Close(ctx)

	var results []schema.Nodes
	if err = cursor.All(ctx, &results); err != nil {
		HandleError(err)
	}
	for _, r := range results {
		fmt.Printf("\nSTART: ['compute_nodes' DUMMY DATA TEST] ~~~~~~~\nID: %v\nCollection: 'compute_nodes'\nDatabase: 'nodes_db'\nDocument: %v\nEND ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n",
			nodeInsertResult.InsertedID, r)
	}

}
