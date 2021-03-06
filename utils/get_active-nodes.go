package utils

import (
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func Get_ActiveNodes() []schema.Nodes {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	nodesDb := client.Database("nodes_db")
	computeNodes := nodesDb.Collection("compute_nodes")

	cursor, err := computeNodes.Find(ctx, bson.M{"active": true})
	HandleError(err)

	defer cursor.Close(ctx)

	var results []schema.Nodes

	if err = cursor.All(ctx, &results); err != nil {
		HandleError(err)
	}

    return results
}
