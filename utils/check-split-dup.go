package utils

import (
	"github.com/iktefish/binary-helix/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckSplitDup() bool {
	client, ctx := ConnectDb()
	defer client.Disconnect(ctx)

	/* Catch slices_db database and slices collection */
	slicesDb := client.Database("slices_db")
	slices := slicesDb.Collection("slices")

	targetIp := ""
	content := ""
	task := ""

	/* NOTE: This will be taken from Carrier() */
	RegisterSlice := schema.Slices{
		SplitOrder: 1,
		Content:    content,
		AnalysisArt: schema.Analysis{
			Task:         task,
			TargetIP:     targetIp,
			Completed:    false,
			Paid:         false,
			UnitOutput:   "",
			MergedOutput: "",
		},
	}

	/* Check for duplicate slice_order */
	cursor, err := slices.Find(ctx, bson.M{"slice_order": RegisterSlice.SplitOrder})
	HandleError(err)
	defer cursor.Close(ctx)
	var results []schema.Slices
	if err = cursor.All(ctx, &results); err != nil {
		HandleError(err)
	}

	if len(results) > 0 {
		return true
	} else {
		return false
	}
}
