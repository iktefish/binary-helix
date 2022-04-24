package utils

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func Admin_EchoDbs() {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	HandleError(err)

	fmt.Println("Currently active databases ~~> ", databases)
}
