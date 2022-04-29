package utils

import (
	"fmt"
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func Admin_DummyInSlices() {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	slicesDb := client.Database("slices_db")
	slices := slicesDb.Collection("slices")

	dummy := schema.Slices{
		ComputationId: "DUMMY",
		SplitOrder:    1,
		Content:       "DUMMY",
		AnalysisArt: schema.Analysis{
			Task:          "DUMMY",
			TargetIP_Port: "DUMMY",
			Completed:     false,
			Paid:          false,
			// UnitOutput:    "DUMMY",
			// MergedOutput:  "DUMMY",
		},
        MergedOutput: []string{"DUMMY"},
	}

	/* Insert dummy in slices collection */
	sliceInsertResult, err := slices.InsertOne(ctx, dummy)
	HandleError(err)

	/* Query slices collection */
	cursor, err := slices.Find(ctx, bson.M{"_id": sliceInsertResult.InsertedID})
	HandleError(err)

	defer cursor.Close(ctx)

	var results []schema.Slices
	if err = cursor.All(ctx, &results); err != nil {
		HandleError(err)
	}
	for _, r := range results {
		fmt.Printf("\nSTART: ['slices' DUMMY DATA TEST] ~~~~~~~~~~~~~~\nID: %v\nCollection: 'slices'\nDatabase: 'slices_db'\nDocument: %v\nEND ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n",
			sliceInsertResult.InsertedID, r)
	}

}
