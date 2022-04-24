package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ConnectToDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	CheckError(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	CheckError(err)

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	CheckError(err)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	CheckError(err)

	fmt.Println("databases ~~> ", databases)
}
