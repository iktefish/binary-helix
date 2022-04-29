package workers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func PayResolver() {

	client, ctx := utils.ConnectDb()
	defer client.Disconnect(ctx)

	/* Catch slices_db database and slices collection */
	slicesDb := client.Database("slices_db")
	slices := slicesDb.Collection("slices")

	// targetIp := ""
	content := ""
	// task := ""

	/* NOTE: This will be taken from Carrier() */
	RegisterSlice := schema.Slices{
		ComputationId: uuid.New().String(),
		SplitOrder:    1,
		Content:       content,
		AnalysisArt: schema.Analysis{
			Task:          "exact-match",
            TargetIP_Port: "127.0.0.1:4040",
			Completed:     false,
			Paid:          false,
			// UnitOutput:    "",
			// MergedOutput:  "",
			// Task:          task,
			// TargetIP_Port: targetIp,
			// Completed:     false,
			// Paid:          false,
			// UnitOutput:    "",
			// MergedOutput:  "",
		},
	}

	/* Check for duplicate slice_order */
	cursor, err := slices.Find(ctx, bson.M{"slice_order": RegisterSlice.SplitOrder})
    // cursor, err := slices.Find(ctx, bson.M{"analysis_art": RegisterSlice.AnalysisArt})
	utils.HandleError(err)
	defer cursor.Close(ctx)
	var results []schema.Slices
	// var results []schema.Analysis
	if err = cursor.All(ctx, &results); err != nil {
		utils.HandleError(err)
	}

    fmt.Println(results)
    pushToDB(results)

	// if len(results) > 0 {
	// 	return true
	// } else {
	// 	return false
	// }

}

func pushToDB(slices []schema.Slices) {
    
}
