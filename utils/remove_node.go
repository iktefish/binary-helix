package utils

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func Unregister_Node(node_name string) {
	client, ctx := ConnectDb()
	db := client.Database("nodes_db")

	result, err := db.Collection("compute_nodes").DeleteOne(ctx, bson.M{"node_name": node_name})
	HandleError(err)

	fmt.Printf("\nSUCCESS: %v documents deleted from database '%v'!\n", result.DeletedCount, node_name)
	fmt.Println()
	defer client.Disconnect(ctx)
}
