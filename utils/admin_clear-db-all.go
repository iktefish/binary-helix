package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

func Admin_ClearDbAll(s string) {
	var wg sync.WaitGroup
	wg.Add(1)

	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	if s == "nodes_db" {
		go clearDb(s, client, ctx, &wg)
	}

	if s == "slices_db" {
		go clearDb(s, client, ctx, &wg)
	}

	if s == "bench_db" {
		go clearDb(s, client, ctx, &wg)
	}

	if s == "all" {
		go clearDb(s, client, ctx, &wg)
	}

	if s != "nodes_db" && s != "slices_db" && s != "bench_db" && s != "all" {
		fmt.Printf("FAIL: The database '%v' does not exist!\n", s)
		wg.Done()
	}

	wg.Wait()
}

func clearDb(s string, c *mongo.Client, ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		if s == "nodes_db" {
            db := c.Database(s)
            colName := "compute_nodes"
			result, err := db.Collection(colName).DeleteMany(ctx, bson.M{})
			HandleError(err)
			fmt.Printf("SUCCESS: `%v` documents deleted from database '%v'!\n", result.DeletedCount, s)
		}

		if s == "slices_db" {
            db := c.Database(s)
            colName := "slices"
			result, err := db.Collection(colName).DeleteMany(ctx, bson.M{})
			HandleError(err)
			fmt.Printf("SUCCESS: `%v` documents deleted from database '%v'!\n", result.DeletedCount, s)
		}

		if s == "bench_db" {
            db := c.Database(s)
            colName := "benchmarks"
			result, err := db.Collection(colName).DeleteMany(ctx, bson.M{})
			HandleError(err)
			fmt.Printf("SUCCESS: `%v` documents deleted from database '%v'!\n", result.DeletedCount, s)
		}

		if s == "all" {
            totalDeleted := 0

            db := c.Database("nodes_db")
            colName := "compute_nodes"
			result, err := db.Collection(colName).DeleteMany(ctx, bson.M{})
			HandleError(err)
            totalDeleted += int(result.DeletedCount)

			db = c.Database("slices_db")
			colName = "slices"
			result, err = db.Collection(colName).DeleteMany(ctx, bson.M{})
			HandleError(err)
            totalDeleted += int(result.DeletedCount)

			db = c.Database("bench_db")
			colName = "benchmarks"
			result, err = db.Collection(colName).DeleteMany(ctx, bson.M{})
			HandleError(err)
            totalDeleted += int(result.DeletedCount)

			fmt.Printf("SUCCESS: `%v` documents deleted from '%v' databases!\n", totalDeleted, s)
		}
}
