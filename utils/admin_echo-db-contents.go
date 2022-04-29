package utils

import (
	"context"
	"fmt"
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

func Admin_EchoDbContents(s string, itemCount *int, mergedICount *int) {
	var wg sync.WaitGroup
	wg.Add(1)

	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	if s == "nodes_db" {
		catchAndEcho(s, client, ctx, &wg, itemCount, mergedICount)
	}

	if s == "slices_db" {
		catchAndEcho(s, client, ctx, &wg, itemCount, mergedICount)
	}

	if s == "bench_db" {
		catchAndEcho(s, client, ctx, &wg, itemCount, mergedICount)
	}

	if s != "nodes_db" && s != "slices_db" && s != "bench_db" {
		fmt.Printf("FAIL:\t The database '%v' does not exist!\n", s)
		wg.Done()
	}

	// nodeInsertResult, err := computeNodes.InsertOne(ctx, schema.Nodes{
	// 	NodeName:                "binary-helix_c1",
	// 	TargetIP:                "172.17.0.2",
	// 	TotalCreditAttained:     0,
	// 	TaskCompletionFrequency: 0,
	// })
	// HandleError(err)
	// fmt.Println("Inserted ~~~> ", nodeInsertResult.InsertedID)
	//
	// /* Query computeDB database */
	// cursor, err := computeNodes.Find(ctx, bson.M{})
	// HandleError(err)
	// defer cursor.Close(ctx)
	// var results []schema.Nodes
	// if err = cursor.All(ctx, &results); err != nil {
	// 	HandleError(err)
	// }
	// for _, r := range results {
	// 	fmt.Println(r)
	// }

	wg.Wait()
}

func catchAndEcho(s string, c *mongo.Client, ctx context.Context, wg *sync.WaitGroup, itemCount *int, mergerICount *int) {
	/* Ready/Catch `s` database and `colName` collection */
	db := c.Database(s)
	colName := ""

	if s == "nodes_db" {
		colName = "compute_nodes"
	}
	if s == "slices_db" {
		colName = "slices"
	}
	if s == "bench_db" {
		colName = "benchmarks"
	}

	collection := db.Collection(colName)

	go func() {
		defer wg.Done()
		cursor, err := collection.Find(ctx, bson.M{})
		HandleError(err)

		defer cursor.Close(ctx)

		if s == "nodes_db" {
			var results []schema.Nodes
			if err = cursor.All(ctx, &results); err != nil {
				HandleError(err)
			}

			if len(results) > 0 {
				for i, r := range results {
					fmt.Printf("%s ~~>\t[%v]\t%v\n", s, i+1, r)
					*itemCount += 1
				}
			} else {
				fmt.Printf("SUCCESS:\t The database '%v' exists but is empty!\n", s)
			}
		}

		if s == "slices_db" {
			var results []schema.Slices
			if err = cursor.All(ctx, &results); err != nil {
				HandleError(err)
			}

			if len(results) > 0 {
				for i, r := range results {
					if len(r.MergedOutput) == 0 {
						*mergerICount += 1
					}
					fmt.Printf("%s ~~>\t[%v]\t%v\n", s, i+1, r)
					*itemCount += 1
				}
			} else {
				fmt.Printf("SUCCESS:\t The database '%v' exists but is empty!\n", s)
			}
		}

		if s == "bench_db" {
			var results []schema.Bench
			if err = cursor.All(ctx, &results); err != nil {
				HandleError(err)
			}

			if len(results) > 0 {
				for i, r := range results {
					fmt.Printf("%s ~~>\t[%v]\t%v\n", s, i+1, r)
					*itemCount += 1
				}
			} else {
				fmt.Printf("SUCCESS:\t The database '%v' exists but is empty!\n", s)
			}
		}

	}()
}
