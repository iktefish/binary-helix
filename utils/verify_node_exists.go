package utils

import (
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func Verify_Node_Exists(node_name string) bool {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	/* Catch nodesDb database and compute_nodes collection */
	nodesDb := client.Database("nodes_db")
	computeNodes := nodesDb.Collection("compute_nodes")

	// Check if node_name exists
	cursor, err := computeNodes.Find(ctx, bson.M{"node_name": node_name})
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
