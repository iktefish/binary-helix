package utils

import (
	"fmt"
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func Admin_DummyInBenchmarks() {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	benchDb := client.Database("bench_db")
	benchmarks := benchDb.Collection("benchmarks")

	dummy := schema.Bench{
		NodeName:      "DUMMY",
		TargetIP_Port:      "DUMMY",
		NetSpeedToS:   0,
		NetSpeedFromS: 0,
		RamTotal:      0,
		RamUsed:       0,
		RamCached:     0,
		RamFree:       0,
		CpuUser:       0,
		CpuSystem:     0,
		CpuIdle:       0,
		TimeStamp:     primitive.NewObjectIDFromTimestamp(time.Now()),
	}

	/* Insert dummy in slices collection */
	benchmarkInsertResult, err := benchmarks.InsertOne(ctx, dummy)
	HandleError(err)

	/* Query slices collection */
	cursor, err := benchmarks.Find(ctx, bson.M{"_id": benchmarkInsertResult.InsertedID})
	HandleError(err)

	defer cursor.Close(ctx)

	var results []schema.Bench
	if err = cursor.All(ctx, &results); err != nil {
		HandleError(err)
	}
	for _, r := range results {
        fmt.Printf("\nSTART: ['benchmarks' DUMMY DATA TEST] ~~~~~~~~~~\nID: %v\nCollection: 'benchmarks'\nDatabase: 'bench_db'\nDocument: %v\nEND ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n",
			benchmarkInsertResult.InsertedID, r)
	}

}
